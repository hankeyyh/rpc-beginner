package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("conn to server %s\n", conn.RemoteAddr())

	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)
	for {
		content, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		content = strings.TrimSpace(content)
		nsend, err := conn.Write([]byte(content))
		if err != nil {
			log.Println(err)
		}
		log.Printf("send %d bytes\n", nsend)
		if nsend == 0 {
			continue
		}

		// recv from server
		nrecv, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
		}
		recvStr := string(buf[:nrecv])
		log.Printf("[server]:%q (%d bytes)\n", recvStr, nrecv)
	}
}
