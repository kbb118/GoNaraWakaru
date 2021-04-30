package main

import (
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	file, err := os.Create("out.html")
	if err != nil {
		panic(err)
	}

	// 自分でhttpリクエストを書く
	//io.WriteString(conn, "GET / HTTP/1.0\r\nHost: example.com\r\n\r\n")

	// http.NewRequestを使う
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		panic(err)
	}
	req.Write(conn)

	io.Copy(file, conn)
}
