package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8080")
	// new request (method, url, body)
	request, _ := http.NewRequest("GET", "http://localhost:8080/about.html", nil)
	request.Write(conn)
	response, _ := http.ReadResponse(bufio.NewReader(conn), request)
	dump, _ := httputil.DumpResponse(response, true)
	fmt.Println(string(dump))
}
