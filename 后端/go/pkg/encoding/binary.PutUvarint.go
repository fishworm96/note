package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	buf := make([]byte, binary.MaxVarintLen64)
	for _, x := range []uint64{1, 2, 127, 128, 255, 256} {
		n := binary.PutUvarint(buf, x)
		fmt.Printf("%x\n", buf[:n])
		// 01
		// 02
		// 7f
		// 8001
		// ff01
		// 8002
	}
}