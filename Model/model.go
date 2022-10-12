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

type Model interface {
	GetUserName() Userlist
	GetUserToken(string) Userlist
	CreateMessage(string, string, string)
	GetMessage(string, string) Messagelist
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

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"user_name"`
	Token string `db:"token"`
}

type Userlist []User

type Message struct {
	ID          int    `db:"id"`
	MessageFrom string `db:"message_from"`
	MessageTo   string `db:"message_to"`
	Message     string `db:"message"`
}

type Messagelist []Message

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

func (model *model) GetUserName() Userlist {
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

func (model *model) GetUserToken(name string) Userlist {

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

func (model *model) CreateMessage(message string, from string, to string) {

	_, err := model.conn.NamedExec(`INSERT INTO message (message,message_from,message_to,date) VALUES (:message,:message_from,:message_to,CURRENT_TIMESTAMP)`,
		map[string]interface{}{
			"message":      message,
			"message_from": from,
			"message_to":   to,
		})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Insert message!")

}

func (model *model) GetMessage(message_id string, message_to string) Messagelist {
	var messagelist Messagelist
	var message Message

	//SELECTを実行。db.Queryの代わりにdb.Queryxを使う。
	rows, err := model.conn.Queryx(fmt.Sprintf("SELECT id,message_to,message_from,message FROM message where id > '%s' and message_to = '%s'", message_id, message_to))
	if err != nil {
		log.Fatal(err)
	}

	// ユーザのTokenを取得
	messagelist = PackData(rows, message, messagelist)

	return messagelist
}
