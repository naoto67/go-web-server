package main

import (
	"fmt"
	"log"
	"net"
)

const (
	BLANK = " "
)

func main() {

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal("fail to connect tcp://localhost:8080/")
	}
	// 終了の通知のためのチャネル
	ch := make(chan struct{})
	go func(end chan<- struct{}) {
		defer conn.Close()
		conn.Write([]byte("GET /index.html HTTP/1.1" + BLANK))
		conn.Write([]byte("Host: localhost:8080" + BLANK))
		conn.Write([]byte("User-Agent: client.go "))
		conn.Write([]byte("0"))
		buf := make([]byte, 1024)
		for {
			n, _ := conn.Read(buf)
			if n == 0 {
				break
			}
			fmt.Println(string(buf))
		}
		//　終了を伝える
		close(end)
	}(ch)

	// ch が受診するのを待ち、通知を受け取ると終了
	for {
		select {
		case <-ch:
			return
		}
	}
}
