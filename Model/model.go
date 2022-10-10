package model

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type model struct {
	conn *sqlx.DB
}

// 引っ張ってきたデータを当てはめる構造体を用意。
// その際、バッククオート（`）で、どのカラムと紐づけるのかを明示する。
type User struct {
	ID    int    `db:"id"`
	Name  string `db:"user_name"`
	Token string `db:"token"`
}

type Userlist []User

func PackData[T comparable](rows *sqlx.Rows, obj T, objlist []T) []T {
	for rows.Next() {
		//rows.Scanの代わりにrows.StructScanを使う
		err := rows.StructScan(&obj)
		if err != nil {
			log.Fatal(err)
		}
		objlist = append(objlist, obj)
	}
	return objlist
}

type Model interface {
	GetUserName() []User
	GetUserToken(string) []User
	CreateMessage(string) *sqlx.Rows
}

func NewModel() Model {
	//Mysqlに接続。sql.Openの代わりにsqlx.Openを使う。
	//ドライバ名、データソース名を引数に渡す
	conn, err := sqlx.Open("sqlite3", "./mobus_db.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	return &model{conn}
}

func (model *model) GetUserName() []User {
	var userlist Userlist
	var user User

	//SELECTを実行。db.Queryの代わりにdb.Queryxを使う。
	rows, err := model.conn.Queryx("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}

	// ユーザのTokenを取得
	userlist = PackData(rows, user, userlist)

	return userlist
}

func (model *model) GetUserToken(name string) []User {

	var userlist Userlist
	var user User

	//SELECTを実行。db.Queryの代わりにdb.Queryxを使う。
	rows, err := model.conn.Queryx(fmt.Sprintf("SELECT * FROM user where user_name = '%s'", name))
	if err != nil {
		log.Fatal(err)
	}

	// ユーザのTokenを取得
	userlist = PackData(rows, user, userlist)

	return userlist
}

func (model *model) CreateMessage(name string) *sqlx.Rows {

	//SELECTを実行。db.Queryの代わりにdb.Queryxを使う。
	rows, err := model.conn.Queryx(fmt.Sprintf("SELECT * FROM user where user_name = '%s'", name))
	if err != nil {
		log.Fatal(err)
	}
	return rows
}
