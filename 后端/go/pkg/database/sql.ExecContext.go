package main

import (
	"context"
	"database/sql"
	"log"
)

var (
	ctx context.Context
	db *sql.DB
)

func main() {
	// *DB是一个连接池。调用 Conn 可保留一个连接，以供专用。
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if := 41
	result, err := conn.ExecContext(ctx, `UPDATE balances SET balance = balance  + 10 WHERE user_id = ?;`, id)
	if err != nil {
		log.Fatal(err)
	}
	
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("expected single row affected, got %d rows affected", rows)
	}
}