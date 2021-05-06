package main

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Diposition", "attachment; filename=hoge.zip")
	source := "hello, from zip!"

	/*
		source --- zip --- http.Response
			   \__ stdout
	*/
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()
	helloWriter, err := zipWriter.Create("hello.txt")
	if err != nil {
		panic(err)
	}
	mulwriter := io.MultiWriter(os.Stdout, helloWriter)
	io.Copy(mulwriter, strings.NewReader(source))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
