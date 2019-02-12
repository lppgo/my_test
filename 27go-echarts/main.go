package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chenjiandongx/go-echarts/charts"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	items := []string{"java", "c", "c#", "js", "python", "go"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"})
	bar.AddXAxis(items).
		AddYAxis("2017年", []int{50, 45, 32, 38, 46, 29}).
		AddYAxis("2018年", []int{46, 30, 30, 40, 41, 35})
	f, err := os.Create("bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(w, f)
}
