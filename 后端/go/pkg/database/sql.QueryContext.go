package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

var (
	ctx context.Context
	db *sql.DB
)

func main() {
	age := 27
	rows, err := db.QueryContext(ctx, "SELECT name FROM users WHERE age = ?", age)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	names := make([]string, 0)

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			// 检查是否有扫描错误。
			// 查询行将通过延迟关闭。
			log.Fatal(err)
		}
		names = append(names, name)
	}

	// 如果正在写入数据库，确保检查驱动程序可能返回的关闭错误。查询可能会遇到自动提交错误，并被迫回滚更改。
	err := rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Rows.Err 将报告 Rows.Scan 遇到的最后一个错误。
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s are %d years old", strings.Join(name, ", ", age))
}