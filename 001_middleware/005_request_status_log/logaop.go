package middlewares
import (
   "bytes"
   "github.com/gin-gonic/gin"
   log "github.com/sirupsen/logrus"
   "time"
)

type responseBodyWriter struct {
   gin.ResponseWriter
   body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
   r.body.Write(b)
   return r.ResponseWriter.Write(b)
}

/**
请求之前
*/
func LogAopReq() func(c *gin.Context) {
   return func(c *gin.Context) {

      //设置日志格式
      log.SetFormatter(&log.JSONFormatter{
         TimestampFormat: "2006-01-02 15:04:05",
      })
      // 开始时间
      startTime := time.Now()

      w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
      c.Writer = w

      // 处理请求
      c.Next()

      // 结束时间
      endTime := time.Now()

      // 执行时间
      latencyTime := endTime.Sub(startTime)

      // 请求方式
      reqMethod := c.Request.Method

      // 请求路由
      reqUri := c.Request.RequestURI

      // 状态码
      statusCode := c.Writer.Status()

      // 请求IP
      clientIP := c.Request.Host

      log.WithFields(log.Fields{
         "status_code":  statusCode,
         "latency_time": latencyTime,
         "client_ip":    clientIP,
         "req_method":   reqMethod,
         "req_uri":      reqUri,
         "response":     w.body.String(),
      }).Info()
   }
}