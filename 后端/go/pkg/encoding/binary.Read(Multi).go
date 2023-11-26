package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40, 0xff, 0x01, 0x02, 0x03, 0xbe, 0xef}
	r := bytes.NewReader(b)

	var data struct {
		PI float64
		Uate uint8
		Mine [3]byte
		Too uint16
	}

	if err := binary.Read(r, binary.LittleEndian, &data); err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Println(data.PI) // 3.141592653589793
	fmt.Println(data.Uate) // 255
	fmt.Println("% x\n", data.Mine) // 01 02 03
	fmt.Println(data.Too) // 61374
}