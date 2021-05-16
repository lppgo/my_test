package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 优雅的关闭服务
var (
	listenAddr string
	// exit       = make(chan struct{}, 1)
)

func main() {
	flag.StringVar(&listenAddr, "listen-addr", ":5000", "server listen address")
	flag.Parse()

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	server := newWebserver(logger)
	go gracefullShutdown(server, logger)

	logger.Println("Server is ready to handle requests at", listenAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	// <-exit
	logger.Println("Server stopped")
}

func gracefullShutdown(server *http.Server, logger *log.Logger) {
	logger.Println("gracefullShutdown is lestening ...")

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// graceful shutdown
	for {
		select {
		case <-ctx.Done(): //收到os.Signal退出信号
			logger.Println("gracefullShutdown is received os.Signal ...")
			goto jump
		default:
		}
	}
jump:
	server.SetKeepAlivesEnabled(false)

	if err := server.Shutdown(ctx); err != nil {
		if err.Error() != context.Canceled.Error() {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
	}
	logger.Println("gracefullShutdown is shutdown ...")
	// close(exit)
}

func newWebserver(logger *log.Logger) *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return &http.Server{
		Addr:         listenAddr,
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
