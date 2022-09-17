package controller

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"local.package/model"
	"local.package/requests"
)

type controller struct {
	model model.Model
}

type Controller interface {
	GetUserName()
	GetUserToken()
	SendMessage(requests.SendMessageRequest)
}

// 引っ張ってきたデータを当てはめる構造体を用意。
// その際、バッククオート（`）で、どのカラムと紐づけるのかを明示する。
type User struct {
	ID    int    `db:"id"`
	Name  string `db:"user_name"`
	Token int    `db:"token"`
}

type Userlist []User

func NewController(model model.Model) Controller {
	return &controller{model}
}

func PackData[T comparable](rows *sqlx.Rows, obj T) {
	type ObjList []T
	var objlist ObjList

	for rows.Next() {
		//rows.Scanの代わりにrows.StructScanを使う
		err := rows.StructScan(&obj)
		if err != nil {
			log.Fatal(err)
		}
		objlist = append(objlist, obj)
	}
}

func (controller *controller) GetUserName() {

	var userlist Userlist
	var user User

	rows := controller.model.GetUserName()

	for rows.Next() {
		//rows.Scanの代わりにrows.StructScanを使う
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
		}
		userlist = append(userlist, user)
	}
	fmt.Println("Called Controller!!")
	fmt.Println(userlist[0].Name)
}

func (controller *controller) GetUserToken() {
	var userlist Userlist
	var user User

	rows := controller.model.GetUserToken()

	for rows.Next() {
		//rows.Scanの代わりにrows.StructScanを使う
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
		}
		userlist = append(userlist, user)
	}
	fmt.Println("Called Controller!!")
	fmt.Println(userlist[0].Token)
}

func (controller *controller) SendMessage(SendMessageJson requests.SendMessageRequest) {

	var userlist Userlist
	var user User

	rows := controller.model.GetUserToken()

	// fmt.Printf("Message Info: %v",SendMessageJson)

	for rows.Next() {
		//rows.Scanの代わりにrows.StructScanを使う
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
		}
		userlist = append(userlist, user)
	}
	fmt.Println("Called Controller!!")
	fmt.Printf("User Token:%v\n", userlist[0].Token)
	// username := userlist[0].Name
}
