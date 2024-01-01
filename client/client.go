package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"rpc-beginner/proto"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("conn to server %s\n", conn.RemoteAddr())

	reader := bufio.NewReader(os.Stdin)
	codec := proto.NewContentGobCodec(conn)
	seq := 0

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}

		req := proto.NewContent(strings.TrimSpace(msg), seq)
		seq++
		err = codec.Encode(req)
		if err != nil {
			if err == io.EOF {
				log.Println("conn closed by server")
				break
			}
			log.Println("codec.Encode error: ", err)
		}

		// recv from server
		var rsp proto.Content
		err = codec.Decode(&rsp)
		if err != nil {
			if err == io.EOF {
				log.Println("conn closed by server")
				break
			}
			log.Println("codec.Decode error: ", err)
		}
		log.Printf("[server] msg: %q, seq: %d\n", rsp.Msg, rsp.Seq)
	}
}
