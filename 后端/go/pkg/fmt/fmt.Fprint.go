package main

import (
	"fmt"
	"os"
)

func main() {
	const name, age = "Kim", 22
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")

	// Fprint 的 n 和 err 返回值是底层 io.Writer 返回的值。
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err) // Kim is 22 years old.
	}
	fmt.Print(n, " bytes written.\n") // 21 bytes written.

}
