package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const (
	BYTE_SIZE = 1024
	ROOT      = "."
	BLANK     = "\n\r"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("fail to listen tcp 8080")
	}
	fmt.Println("listening on 8080")
	buf := make([]byte, BYTE_SIZE)
	for {
		conn, _ := listen.Accept()
		go func() {
			defer conn.Close()
			n, _ := conn.Read(buf)
			str := string(buf)
			path := strings.Split(str, " ")[1]
			fmt.Println(string(buf[:n]))

			conn.Write([]byte("HTTP/1.1 200 OK" + BLANK))
			conn.Write([]byte("Date: " + time.Now().String() + BLANK))
			conn.Write([]byte("" + path + BLANK))
			conn.Write([]byte("Server: superserver.go" + BLANK))
			conn.Write([]byte("Connection: close" + BLANK))
			conn.Write([]byte("Content-type: text/html" + BLANK))
			conn.Write([]byte("" + BLANK))

			file_path := ROOT + path

			file, err := os.Open(file_path)
			if err != nil {
				log.Println(err)
				log.Println("fail to open file" + path)
			}
			defer file.Close()
			for {
				n, err := file.Read(buf)
				if n == 0 {
					break
				}
				if err != nil {
					break
				}
				conn.Write(buf)
			}
		}()
	}
}
