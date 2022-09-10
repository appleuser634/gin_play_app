package main

import (
    "fmt"
    "log"

    "github.com/jmoiron/sqlx"
    _ "github.com/mattn/go-sqlite3"

    "github.com/gin-gonic/gin"
)

//引っ張ってきたデータを当てはめる構造体を用意。
//その際、バッククオート（`）で、どのカラムと紐づけるのかを明示する。
type User struct {
    ID   int    `db:"id"`
    Name string `db:"name"`
    Age  int    `db:"age"`
}

type Userlist []User

func main() {
    r := gin.Default()

    //Userデータ一件一件を格納する配列Userlistを、Userlist型で用意
    var userlist Userlist

    //Mysqlに接続。sql.Openの代わりにsqlx.Openを使う。
    //ドライバ名、データソース名を引数に渡す
    db, err := sqlx.Open("sqlite3", "./test.sqlite")
    if err != nil {
        log.Fatal(err)
    }


    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    
    r.GET("/getUserName", func(c *gin.Context) {
        //SELECTを実行。db.Queryの代わりにdb.Queryxを使う。
        rows, err := db.Queryx("SELECT * FROM test")
        if err != nil {
            log.Fatal(err)
        }

        var user User
        for rows.Next() {

            //rows.Scanの代わりにrows.StructScanを使う
            err := rows.StructScan(&user)
            if err != nil {
                log.Fatal(err)
            }
            userlist = append(userlist, user)
        }

        fmt.Println(userlist[0].Name)
        username := userlist[0].Name
        //[{1 yamada 25} {2 suzuki 28}]
        c.JSON(200, gin.H{
            "message": username,
        })
    })


    r.Run(":3000")
}
