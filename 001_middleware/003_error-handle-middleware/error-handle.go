package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

}

//
type ErrorHandle func(request http.ResponseWriter, response *http.Request) error

// 优雅的错误处理middleware .
// 通过该middleware 统一的处理panic和error
func GracefulErrorHandle(eh ErrorHandle) func(writer http.ResponseWriter,
	request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 处理panic
		defer func() {
			r := recover()
			if e, ok := r.(error); ok {
				fmt.Print("this is error \n", e)
				http.Error(writer, "please check /", http.StatusNotFound)
			} else {
				panic(e)
			}
		}()
		// 处理 handleFunc返回的error
		err := eh(writer, request)
		if err != nil {
			switch { // 各种错误类型的判断
			case os.IsNotExist(err):
				http.Error(writer, http.StatusText(http.StatusNotFound),
					http.StatusNotFound)
			default:
				http.Error(writer, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}
	}
}
