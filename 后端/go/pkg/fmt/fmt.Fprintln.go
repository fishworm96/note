package main

import (
	"fmt"
	"os"
)

func main() {
	const name, age = "Kim", 22
	n, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old.")

	// Fprintln 的 n 和 err 返回值是底层 io.Writer 返回的值。
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err) // Kim is 22 years old.
	}
	fmt.Println(n, "bytes written.") // 21 bytes written.
}
