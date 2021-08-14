package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db1")
	if err != nil {
		log.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	res, _ := db.Query("SHOW TABLES")

	var table string

	for res.Next() {
		res.Scan(&table)
		fmt.Println(table)
	}
}
