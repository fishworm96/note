package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	buf := make([]byte, binary.MaxVarintLen64)

	for _, x := range []int64{-65, -64, -2, -1, 0, 1, 2, 63, 64} {
		n := binary.PutVarint(buf, x)
		fmt.Printf("%x\n", buf[:n])
	}
	// 8101
	// 7f
	// 03
	// 01
	// 00
	// 02
	// 04
	// 7e
	// 8001
}