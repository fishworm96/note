package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	buf := new(bytes.Buffer)
	var pi float64 = math.Pi
	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes()) // 18 2d 44 54 fb 21 09 40
}