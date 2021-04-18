package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client) //client上线
	leaving  = make(chan client) //client离线
	messages = make(chan string) //all incoming client messages
)

// broadcaster 广播消息.
func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

// handleConn 处理连接
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + "has online!"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ":" + input.Text()
	}
	// 	NOTE: ignore potential errors from input.Err()

	leaving <- ch
	messages <- who + "has offline!"
	conn.Close()
}

// clientWriter
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE:ignore network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln("listen err:", err.Error())
	}
	log.Println("listened to tcp localhost:8000")

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listener.Accept err:", err.Error())
			continue
		}
		go handleConn(conn)
	}

}
