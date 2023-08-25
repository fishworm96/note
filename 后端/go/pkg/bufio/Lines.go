package main

import (
	"bufio"
	"fmt"
	"os"
)

// 扫描器的最简单用法，将标准输入作为一组行读取。
func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println 会加回最后一个 '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}