package main

import (
	"fmt"
	"os"
)

func main() {
	const name, age = "Kim", 22
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)

	// Fprintf 的 n 和 err 返回值是底层 io.Writer 返回的值。
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err) // Kim is 22 years old.
	}
	fmt.Printf("%d bytes written.\n", n) // 21 bytes written.

}
