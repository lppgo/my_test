package main

import (
	"log"
	"net"
	"time"
)

// TCP Client
func main() {
	// step 1: dial a TCP session to server
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("dial TCP error:%v\n", err)
	}
	defer conn.Close()

	// step 2: write some data to server
	log.Println("TCP session open ...")
	buf := []byte("Hi ,Gopher ,this is a TCP connection !")
	_, err = conn.Write(buf)
	if err != nil {
		log.Fatalf("write data to TCP server error:%v\n", err)
	}

	// step 3: create a goroutine that close TCP session after 10 s
	go CloseConn(conn)

	// step 4: read any response until get an error
	for {
		readBuf := make([]byte, 1024)
		_, err = conn.Read(readBuf)
		if err != nil {
			log.Printf("TCP conn read error:%v\n", err)
			break
		}
		log.Printf("reading data from server: %s\n", string(readBuf))
	}

}

func CloseConn(conn net.Conn) {
	<-time.After(time.Second * 10)
	log.Printf("close TCP conn ...")
	err := conn.Close()
	if err != nil {
		log.Printf("TCP conn close error:%v\n", err)
	}

}
