package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		// logging and Fatal means exit.
		log.Fatal("fail to listen tcp://localhost:8080")
	}
	fmt.Println("listening on tcp://localhost:8080")
	buf := make([]byte, 1024)
	for {
		conn, _ := listen.Accept()
		go func() {
			// defer は 逆さまの順番に実行される

			// conn.Close() でコネクションを終了しないと、
			// リクエスト投げたままクライアントがレスポンス待ちの状態になる
			defer conn.Close()
			defer conn.Write([]byte("コネクション終了\n"))

			// n は バイト数, buf に 読み込んだ内容が入る
			// n, _ := conn.Read(buf)

			// ex. buf

			//  GET / HTTP/1.1
			//  Host: localhost:8080
			//  User-Agent: curl/7.54.0
			//  Accept: */*
			//  content-type: application/json
			conn.Read(buf)
			fmt.Println(string(buf))
		}()
	}
}
