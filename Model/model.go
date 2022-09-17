package model

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type model struct {
	conn *sqlx.DB
}

type Model interface {
	GetUserName() *sqlx.Rows
	GetUserToken() *sqlx.Rows
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

func (model *model) GetUserName() *sqlx.Rows {

	//SELECTを実行。db.Queryの代わりにdb.Queryxを使う。
	rows, err := model.conn.Queryx("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func (model *model) GetUserToken() *sqlx.Rows {

	//SELECTを実行。db.Queryの代わりにdb.Queryxを使う。
	rows, err := model.conn.Queryx("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	return rows
}
