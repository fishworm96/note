package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// 使用具有自定义拆分功能（通过包装ScanWords构建）的扫描仪进行验证 32-位十进制输入。
func main()  {
	// 人工输入源
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 通过包装现有 ScanWords 函数创建自定义拆分函数
	split := func (data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	// 为扫描操作设置拆分函数。
	scanner.Split(split)
	// 验证输入
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}