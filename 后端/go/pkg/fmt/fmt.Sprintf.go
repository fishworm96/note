package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age) // Kim is 22 years old.

	io.WriteString(os.Stdout, s) // 为简单起见，忽略误差。

}
