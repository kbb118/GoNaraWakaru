package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func addZip(zipWriter *zip.Writer, name string) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer, err := zipWriter.Create(name)
	io.Copy(writer, file)
}

func main() {
	file, err := os.Create("test.zip")
	if err != nil {
		panic(err)
	}
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	addZip(zipWriter, "hoge.txt")
	addZip(zipWriter, "fuga.txt")

	str := "hogefuga"
	strReader := strings.NewReader(str)
	hogeWriter, err := zipWriter.Create("hogefuga.txt")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(hogeWriter, strReader)
	if err != nil {
		panic(err)
	}
}
