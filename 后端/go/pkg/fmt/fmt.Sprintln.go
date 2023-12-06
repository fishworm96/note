package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const name, age = "Kim", 22
	s := fmt.Sprintln(name, "is", age, "years old.") // Kim is 22 years old.

	io.WriteString(os.Stdout, s) // 为简单起见，忽略误差。

}
