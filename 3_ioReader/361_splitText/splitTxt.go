package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `1行め
2行め
3行め`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		// line には改行文字も込み
		line, err := reader.ReadString('\n')
		fmt.Printf("%v", line)
		if err == io.EOF {
			fmt.Print("\n")
			break
		}
	}

	// Scan の区切り文字デフォルトは改行文字
	// こっちは区切り文字は消される
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%v", scanner.Text())
	}
}
