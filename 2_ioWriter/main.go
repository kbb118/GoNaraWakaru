package main

import (
	"fmt"
)

// インターフェースを定義
type Talker interface {
	Talk()
}

// 構造体を宣言
type Greeter struct {
	name string
}

// 構造体はTalkerインターフェースで定義されるメソッドを持っている
// (Goにはクラスは無く、構造体しかないが構造体がメソッドを持てる。
//  funcキーワードとメソッドのシグネチャの定義の間に「レシーバ」
//  を置くと、構造体にメソッドを定義したことになる。)
func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	// インターフェースの型を持つ変数を宣言
	var talker Talker
	// インターフェースを満たす構造体のポインタは代入できる
	talker = &Greeter{"wozozo"}
	talker.Talk()
}
