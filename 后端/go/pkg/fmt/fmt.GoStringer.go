package main

import (
	"fmt"
)

// 地址包含城市、州和国家。
type Address struct {
	City    string
	State   string
	Country string
}

// 个人有姓名、年龄和地址。
type Person struct {
	Name string
	Age  uint
	Addr *Address
}

// GoString 使 Person 满足 GoStringer 接口，返回值是有效的 Go 代码，可用来重现 Person 结构。
func (p Person) GoString() string {
	if p.Addr != nil {
		return fmt.Sprintf("Person{Name: %q, Age: %d, Addr: &Address{City: %q, State: %q, Country: %q}}", p.Name, int(p.Age), p.Addr.City, p.Addr.State, p.Addr.Country)
	}
	return fmt.Sprintf("Person{Name: %q, Age: %d}", p.Name, int(p.Age))
}

func main() {
	p1 := Person{
		Name: "Warren",
		Age:  31,
		Addr: &Address{
			City:    "Denver",
			State:   "CO",
			Country: "U.S.A.",
		},
	}
	// 如果没有实现 GoString()，`fmt.Printf("%#v", p1)` 的输出将类似于 Person{Name: "Warren", Age:0x1f, Addr:(*main.Address)(0x10448240)} 。
	fmt.Printf("%#v\n", p1) // Person{Name: "Warren", Age: 31, Addr: &Address{City: "Denver", State: "CO", Country: "U.S.A."}}

	p2 := Person{
		Name: "Theia",
		Age:  4,
	}
	// 如果没有实现 GoString()，`fmt.Printf("%#v", p2)` 的输出将类似于 Person{Name: "Theia", Age:0x4, Addr:(*main.Address)(nil)} 。
	fmt.Printf("%#v\n", p2) // Person{Name: "Theia", Age: 4}

}
