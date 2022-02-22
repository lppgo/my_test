package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

func DownLoad(w http.ResponseWriter, r *http.Request) {
	DownLoadFileHandler(w, r, "./readme.md", "./")
}
func main() {
	http.HandleFunc("/", DownLoad)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
func DownLoadFileHandler(w http.ResponseWriter, req *http.Request, fpath string, resPrefix string) {
	defer req.Body.Close()
	err := req.ParseForm()
	//request header
	ranStr := req.Header.Get("Range")
	fmt.Println("Range: " + ranStr)
	//    w.Header().Set("Access-Control-Allow-Origin", "*")
	if err != nil {
		w.Header().Set("Content-Type", " text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if fpath == "" || !strings.HasPrefix(fpath, resPrefix) {
		w.Header().Set("Content-Type", " text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("No file can be downloaded."))
		return
	}
	file, err := os.Open(fpath)
	if err != nil {
		w.Header().Set("Content-Type", " text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
		return
	}
	defer file.Close()
	//设置下载完成以后的名称
	aliasName := path.Base(fpath)
	w.Header().Set("Content-Disposition", "attachment; filename="+aliasName)
	//设置下载的偏移量
	fstat, _ := file.Stat()
	fsize := fstat.Size()
	var spos int64
	if ranStr != "" {
		rs := strings.Split(strings.TrimPrefix(ranStr, "bytes="), ",")[0]
		sePos := strings.Split(rs, "-")
		spStr := sePos[0]
		spos, _ = strconv.ParseInt(spStr, 0, 64)
		file.Seek(spos, 0)
	}
	fmt.Println(spos)
	w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", spos, fsize-1, fsize))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fsize-spos))
	//    w.WriteHeader(http.StatusPartialContent)
	if spos == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusPartialContent)
	}
	wcount, err := io.Copy(w, file)
	if err != nil {
		fmt.Println("io.Copy: ", err.Error())
	}
	if spos+wcount == fsize {
		fmt.Println("remove from cache info.")
		fmt.Println("delete finished file." + fpath)
		file.Close()
		err := os.Remove(fpath)
		if err != nil {
			fmt.Println(err)
		}
	}
}
