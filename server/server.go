package main

import (
	"io"
	"log"
	"net"

	"rpc-beginner/proto"
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
	// buf := make([]byte, 1024)

	codec := proto.NewContentGobCodec(conn)

	for {
		var req proto.Content
		err := codec.Decode(&req)
		if err != nil {
			if err == io.EOF {
				log.Printf("conn closed from %s\n", conn.RemoteAddr())
				break
			}
			log.Println(err)
		}
		log.Printf("[client] msg: %q, seq: %d\n", req.Msg, req.Seq)

		rsp := proto.NewContent("echo " + req.Msg, req.Seq)
		err = codec.Encode(rsp)
		if err != nil {
			if err == io.EOF {
				log.Printf("conn closed from %s\n", conn.RemoteAddr())
				break
			}
			log.Println(err)
		}
	}
}
