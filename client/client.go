package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("conn to server %s\n", conn.RemoteAddr())

	reader := bufio.NewReader(os.Stdin)
	for {
		content, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println(err)
		}
		nbyte, err := conn.Write(content[:len(content) - 1])
		if err != nil {
			log.Println(err)
		}
		log.Printf("send %d bytes\n", nbyte)
	}
}
