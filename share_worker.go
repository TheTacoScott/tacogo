package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const SQL_CREATE_SHARE_TABLE = "create table if not exists 'shares' ('id' INTEGER PRIMARY KEY AUTOINCREMENT,'name' VARCHAR(128) NULL,path VARCHAR(4096) NULL)"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Taco")

	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	res, err := db.Exec(SQL_CREATE_SHARE_TABLE)
	_ = res
	checkErr(err)

	db.Close()
	fmt.Println("Taco2")
}
