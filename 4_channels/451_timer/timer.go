package main

import (
	"time"
)

func main() {
	tchan := time.After(3 * time.Second)
	//time.Sleep(4 * time.Second)
	<-tchan
}
