package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const name, age = "Kim", 22
	s := fmt.Sprint(name, " is ", age, " years old.\n") // Kim is 22 years old.

	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

}
