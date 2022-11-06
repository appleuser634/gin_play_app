package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

func main() {
	type User struct {
		ID    int    `db:"id"`
		Name  string `db:"user_name"`
		Token string `db:"token"`
	}

	ctx := context.Background()

	sqldb, err := sql.Open("sqlite3", "./mobus_db.sqlite")
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	// count, err := db.NewSelect().Model((*User)(nil)).Count(ctx)
	// fmt.Printf("Count:%v", count)

	var user []User
	count, err := db.NewSelect().Model(&user).Limit(20).ScanAndCount(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(user, count)
}
