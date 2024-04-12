package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	url := "postgres://nweznkyg:RAg4JxPS1_QM-AYFd49UOVpyx8UFbn1z@rain.db.elephantsql.com/nweznkyg"
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createTb := `CREATE TABLE IF NOT EXISTS users ( id SERIAL PRIMARY KEY, name TEXT);`
	_, err = db.Exec(createTb)
	if err != nil {
		panic(err)
	}

	row := db.QueryRow("INSERT INTO users(name) VALUES($1) RETURNING id", "AnuchitO")
	var id int
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted ID:", id)

	log.Println("Connected to database")
}
