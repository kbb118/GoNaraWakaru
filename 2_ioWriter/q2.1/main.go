package main

import (
	"fmt"
)

func main() {
	w := 10.0
	h := 20.0
	s := "面積を計算します"
	fmt.Printf("%s\n", s)
	fmt.Printf("幅 %f, 高さ %f の四角形の面積は %f\n", w, h, w*h)
}
