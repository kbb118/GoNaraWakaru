package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	file, err := os.Create("file.bin")
	if err != nil {
		panic(err)
	}
	io.CopyN(file, rand.Reader, 1024)
}
