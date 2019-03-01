package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
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
			defer conn.Close()
			fmt.Printf("リモートアドレスは:%v\n", conn.RemoteAddr())
			conn.SetReadDeadline(time.Now().Add(5 * time.Second))
			request, _ := http.ReadRequest(bufio.NewReaderSize(conn, BYTE_SIZE))
			dump, _ := httputil.DumpRequest(request, true)
			// fmt.Println(request.URL)
			fmt.Printf("%v\n", string(dump))
			body, err := handleRequest(request)
			if err == nil {
				response := http.Response{
					StatusCode: 200,
					ProtoMajor: 1,
					ProtoMinor: 0,
					Body:       ioutil.NopCloser(strings.NewReader(body)),
				}
				response.Write(conn)
			} else {
				response := http.Response{
					StatusCode: 400,
					ProtoMajor: 1,
					ProtoMinor: 0,
				}
				response.Write(conn)
			}
		}()
	}
}

func handleRequest(req *http.Request) (string, error) {
	path := req.URL.String()
	file_path := ROOT + path
	file, err := os.Open(file_path)
	if err != nil {
		log.Println("file not exist")
		return "", err
	}
	defer file.Close()
	body, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("cannot read file.")
	}
	return string(body), err
}
