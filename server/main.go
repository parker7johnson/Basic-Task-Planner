package main

import (
	"fmt"
	controller "server/controller"
	schema "server/db"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sqlx.Open("sqlite3", "./task.db")

	db.Exec(schema.Schema)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully opened first sqlite db: %v\n", db)

	defer db.Close()

	controller.BeginServer(db)

}
