package main

import (
	"io"
	"os"
	"strings"
)

func myIoCopy(dest io.Writer, src io.Reader, length int) {
	lr := io.LimitReader(src, int64(length))
	_, err := io.Copy(dest, lr)
	if err != nil {
		panic(err)
	}
}

func main() {
	str := "Hello, world!"
	myIoCopy(os.Stdout, strings.NewReader(str), 5)
	io.CopyN(os.Stdout, strings.NewReader(str), 5)
}
