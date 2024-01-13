package app

import (
	"arfifa/belajar-golang-restful-api/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "phpmyadmin:blackot@tcp(localhost:3306)/belajar_golang_database_migration")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
