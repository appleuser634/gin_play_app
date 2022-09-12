package Model

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)



//引っ張ってきたデータを当てはめる構造体を用意。
//その際、バッククオート（`）で、どのカラムと紐づけるのかを明示する。
type User struct {
    ID   int    `db:"id"`
    Name string `db:"name"`
    Age  int    `db:"age"`
}

type Userlist []User
func GetUserName() {

	//Mysqlに接続。sql.Openの代わりにsqlx.Openを使う。
	//ドライバ名、データソース名を引数に渡す
	db, err := sqlx.Open("sqlite3", "./test.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	//SELECTを実行。db.Queryの代わりにdb.Queryxを使う。
	rows, err := db.Queryx("SELECT * FROM test")
	if err != nil {
		log.Fatal(err)
	}
	
    //Userデータ一件一件を格納する配列Userlistを、Userlist型で用意
    var userlist Userlist

    var user User

	for rows.Next() {
		//rows.Scanの代わりにrows.StructScanを使う
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
		}
		userlist = append(userlist, user)
	}
		fmt.Println("Called Model!!")
        fmt.Println(userlist[0].Name)
        // username := userlist[0].Name

}
