package middleware

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"git.yeeuu.com/yeeuu/trail/config"
)

// gin 框架dumpbody中间件
var logger *log.Logger
var (
	wrapBytes = []byte("\n")
	nopeBytes = []byte("")
)

type bufWriter struct {
	gin.ResponseWriter
	Writer io.Writer
}

func (bw bufWriter) Write(data []byte) (int, error) {
	bw.Writer.Write(data)
	return bw.ResponseWriter.Write(data)
}

// drainBody reads all of b to memory and then returns two equivalent
// ReadClosers yielding the same bytes.
//
// It returns an error if the initial slurp of all bytes fails. It does not attempt
// to make the returned ReadClosers have identical error-matching behavior.
func drainBody(b io.ReadCloser) (r1 []byte, r2 io.ReadCloser, err error) {
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return buf.Bytes(), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

// DumpReqAndResp dumps request and response. NOTICE: it will be expose some sensitive data.
func DumpReqAndResp() gin.HandlerFunc {
	return func(c *gin.Context) {
		today := time.Now().Format("2006-01-02")
		dumpfile, dumpFileErr := os.OpenFile(config.Get("LOGPATH")+"/dumpbodylog/"+today+"_dumpbody.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
		if dumpFileErr != nil {
			log.Fatalln("Create dumpbody.log file failed:", dumpFileErr.Error())
		}
		defer dumpfile.Close()
		logger = log.New(dumpfile, "", log.LstdFlags|log.Lshortfile)
		var saved []byte
		var err error
		var buf bytes.Buffer
		if c.Request.Body != nil && c.Request.Body != http.NoBody {
			saved, c.Request.Body, err = drainBody(c.Request.Body)
			if err == nil {
				logger.Printf(
					"[%s]%s: %s\n",
					c.Request.Method, c.Request.URL.String(),
					string(bytes.Replace(saved, wrapBytes, nopeBytes, -1)),
				)
			}
		}
		w := bufWriter{c.Writer, &buf}
		c.Writer = w
		c.Next()
		if buf.Len() > 0 {
			if !strings.Contains(c.Writer.Header().Get("Content-Type"), "html") {
				logger.Printf(
					"[%s]%s Resp: %s\n",
					c.Request.Method, c.Request.URL.String(),
					string(bytes.Replace(buf.Bytes(), wrapBytes, nopeBytes, -1)),
				)
			}
		}
	}
}
