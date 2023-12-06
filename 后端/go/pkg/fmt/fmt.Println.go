package main

import (
	"fmt"
)

func main() {
	const name, age = "Kim", 22
	fmt.Println(name, "is", age, "years old.") // Kim is 22 years old.

	// 常规做法是不用担心 Println 返回的任何错误。

}
