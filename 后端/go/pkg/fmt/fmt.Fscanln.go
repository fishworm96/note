package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	s := `dmr 1771 1.61803398875
	ken 271828 3.14159`
	r := strings.NewReader(s)
	var a string
	var b int
	var c float64
	for {
		n, err := fmt.Fscanln(r, &a, &b, &c) // 3: dmr, 1771, 1.618034
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d: %s, %d, %f\n", n, a, b, c) // 3: ken, 271828, 3.141590
	}
}
