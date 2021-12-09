package main

import (
	"io"
	"log"
	"net"
)

// TCP Server
func main() {
	// step1 :
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("listen error:%v", err)
	}
	defer listener.Close()

	// step 2:
	for {
		conn, err := listener.Accept()
		if err == io.EOF {
			continue
		}
		if err != nil {
			log.Printf("accept a new connection is error:%v\n", err)
		}

		// step3 : go to handler conn
		go ConnHandler(conn)

	}
}

func ConnHandler(conn net.Conn) {
	log.Println("TCP session open ...")
	defer conn.Close()

	for {
		// Read from tcp buffer
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err == io.EOF {
			continue
		}
		if err != nil {
			log.Printf("read from TCP error:%v\n", err)
			break
		}
		log.Printf("read from TCP data :%v\n", string(buf))

		// Write data to TCP
		_, err = conn.Write(buf)
		if err == io.EOF {
			continue
		}
		if err != nil {
			log.Printf("write to TCP error:%v\n", err)
			break
		}

	}
}
