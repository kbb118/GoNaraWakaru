package main

import (
	"fmt"
)

func sub(done chan bool) {
	fmt.Println("sub() is finished")
	done <- true
}

func main() {
	fmt.Println("start sub()")
	done := make(chan bool)
	go func() {
		fmt.Println("sub() is finished")
		done <- true
	}()
	<-done

	go sub(done)
	<-done

	fmt.Println("all tasks are finished")
}
