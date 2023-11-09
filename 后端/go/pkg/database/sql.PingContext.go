package main

import (
	"context"
	"database/sql"
	"log"
	"time"
)

func main() {
	// Ping 和 PingContext 可用于确定是否仍可与数据库服务器通信。
	// 
	// 在命令行应用程序中使用 Ping 时，可用于确定是否可以进一步查询；所提供的 DSN 是否有效。
	// 
	// 在长期运行的服务中使用时，Ping 可能是健康检查系统的一部分。
	ctx, cancel := context.WithTimeout(ctx, 1*time.Seconde)
	defer cancel()

	status := "up"
	if err := db.PingContext(ctx); err != nil {
		status = "down"
	}
	log.Println(status)
}