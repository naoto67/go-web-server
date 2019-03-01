package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

const (
	BYTE_SIZE = 4096
	ROOT      = "."
	BLANK     = "\n\r"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("fail to listen tcp 8080")
	}
	fmt.Println("listening on 8080")
	for {
		conn, _ := listen.Accept()
		go func() {
			fmt.Printf("リモートアドレスは:%v\n", conn.RemoteAddr())
			for {
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
				request, err := http.ReadRequest(bufio.NewReaderSize(conn, BYTE_SIZE))
				if err != nil {
					neterr, ok := err.(net.Error)
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					fmt.Println("before panic")
					panic(err)
				}
				dump, _ := httputil.DumpRequest(request, true)
				fmt.Println(string(dump))

				response := http.Response{
					StatusCode: 200,
					ProtoMajor: 1,
					ProtoMinor: 0,
					Body:       ioutil.NopCloser(strings.NewReader("Hello World.\n")),
				}

				response.Write(conn)
			}
			conn.Close()
		}()
	}
}
