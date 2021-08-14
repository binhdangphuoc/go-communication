package config

import (
	"database/sql"
)

func GetMySQLDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	// if there is an error opening the connection, handle it

	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(20)
	//db.SetConnMaxLifetime(time.Second*10)
	return db, err
}
