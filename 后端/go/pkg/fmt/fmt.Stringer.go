package main

import (
	"fmt"
)

// 动物有一个名字和一个年龄，代表一种动物。
type Animal struct {
	Name string
	Age  uint
}

// String 使 Animal 满足 Stringer 接口的要求。
func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func main() {
	a := Animal{
		Name: "Gopher",
		Age:  2,
	}
	fmt.Println(a) // Gopher (2)
}
