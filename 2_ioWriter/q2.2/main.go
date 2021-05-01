package main

/*
import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	mulwriter := io.MultiWriter(file, os.Stdout)
	io.WriteString(mulwriter, "io.MultiWriter example")
}
*/

/*
import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	gwriter := gzip.NewWriter(file)
	gwriter.Header.Name = "test.txt"
	io.WriteString(gwriter, "gzip.Writer example\n")
	gwriter.Close()
}
*/

/*
import (
	"encoding/json"
	"os"
)

func main() {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})
}
*/

import (
	"encoding/csv"
	"os"
)

func main() {
	csvwriter := csv.NewWriter(os.Stdout)
	csvwriter.Write([]string{"hoge", "foo", "bar"})
	csvwriter.Write([]string{"hoge", "foo"})
	csvwriter.Write([]string{"hoge", "foo", "baz"})
	csvwriter.Flush()
}
