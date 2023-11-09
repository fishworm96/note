package main

import (
	"context"
	"database/sql"
	"log"
)

var db *sql.DB

func main() {
	projects := []struct {
		mascot string
		release int
	}{
		{"tux", 1991},
		{"duke", 1996},
		{"gopher", 2009},
		{"moby dock", 2013},
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback() // 如果 tx 已在函数后面提交，回滚将被忽略。

	stmt, err := tx.Prepare("INSERT INTO projects(id, mascot, release, category) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // 准备语句占用服务器资源，使用后应关闭。

	for id, project := range projects {
		if _, err := stmt.Exec(id + 1, project.mascot, project.release, "open source"); err != nil {
			log.Fatal(err)
		}
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}