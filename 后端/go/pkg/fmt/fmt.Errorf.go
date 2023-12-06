package main

import (
	"fmt"
)

// 通过 Errorf 函数，我们可以使用格式化功能创建描述性错误信息。
func main() {
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error()) // user "bueller" (id 17) not found

}
