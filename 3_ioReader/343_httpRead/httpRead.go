package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

func main() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: example.com\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Header)
	defer res.Body.Close()
	fmt.Println(res.Body)
}
