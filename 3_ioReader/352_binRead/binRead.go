package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

func main() {
	rdata := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	binary.Read(bytes.NewReader(rdata), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)

	var j int
	var k int32
	fmt.Printf("sizeof int  : %d\n", unsafe.Sizeof(j))
	fmt.Printf("sizeof int32: %d\n", unsafe.Sizeof(k))
	// -> sizeof int  : 8
	// -> sizeof int32: 4
}
