package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

const SQL_CREATE_SHARE_TABLE = "create table if not exists 'shares' ('id' INTEGER PRIMARY KEY AUTOINCREMENT,'name' VARCHAR(128) NULL,path VARCHAR(4096) NULL)"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func processWalk(path string, info os.FileInfo, err error) error {
	fmt.Println(path, info)
	return nil
}

func main() {
	fmt.Println("Taco")

	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	res, err := db.Exec(SQL_CREATE_SHARE_TABLE)
	_ = res
	checkErr(err)
	filepath.Walk("/home/scott/", processWalk)

	db.Close()
	fmt.Println("Taco2")
}
