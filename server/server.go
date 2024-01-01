package main

import (
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("start server")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		
		go serveConn(conn)
	}
}

func serveConn(conn net.Conn) {
	defer conn.Close()
	log.Printf("conn from %s\n", conn.RemoteAddr())
	
	for {
		buf := make([]byte, 1024)
		nbyte, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Printf("conn closed from %s\n", conn.RemoteAddr())
				break
			}
			log.Println(err)
		}
		log.Printf("[client]: %s (%d bytes)\n", buf, nbyte)
	}
}
