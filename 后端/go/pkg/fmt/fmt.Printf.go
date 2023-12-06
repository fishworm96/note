package main

import (
	"fmt"
)

func main() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age) // Kim is 22 years old.

	// 常规做法是不用担心 Printf 返回的任何错误。

}
