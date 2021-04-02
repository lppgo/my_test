package middlewares

import (
	"net/http"
)

// (1)----------------------------------------------------------
// basic cors
type MyHandler struct {
	handler http.Handler
}

//
func (c *MyHandler) ServerHttp(w http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("origin"); len(origin) > 0 {
		SetHeaders(w, req)
	}

	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	c.handler.ServeHTTP(w, req)
}

// SetHeaders sets the CORS headers
func SetHeaders(w http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("origin"); len(origin) > 0 {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}

	w.Header().Set("Access-Control-Max-Age", "3600")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	// 添加自定义的header
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
}
