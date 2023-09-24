package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 以[]字节的形式返回最近对Scan的调用。
func main()  {
	scanner := bufio.NewScanner(strings.NewReader("gopher"))
	for scanner.Scan() {
		fmt.Println(len(scanner.Bytes()) == 6)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
	}
}