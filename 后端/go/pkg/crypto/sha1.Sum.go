package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	data := []byte("This page intentionally left blank.")
	fmt.Printf("%x", sha1.Sum(data)) // af064923bbf2301596aac4c273ba32178ebc4a96
}