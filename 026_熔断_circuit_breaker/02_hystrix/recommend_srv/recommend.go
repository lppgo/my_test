package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

// 这是一个recommend server
// Recommendation Service
func main() {
	logGoroutines()
	http.HandleFunc("/recommendations", recoHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func logGoroutines() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Printf("\n%v - %v", t, runtime.NumGoroutine())
			}
		}
	}()
}

func recoHandler(w http.ResponseWriter, r *http.Request) {
	a := `{"movies": ["Few Angry Men", "Pride & Prejudice"]}`
	w.Write([]byte(a))
}
