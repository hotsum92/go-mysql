package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test?parseTime=true")
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/test?parseTime=true")

	if err != nil {
		println(fmt.Sprintf("Error: %s", err))
		os.Exit(1)
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO user (name, email) VALUES (?, ?)", "John Doe", "mail")

	if err != nil {
		println(fmt.Sprintf("Error: %s", err))
		os.Exit(1)
	}

	rows, err := db.Query("SELECT * FROM user")

	if err != nil {
		println(fmt.Sprintf("Error: %s", err))
		os.Exit(1)
	}

	for rows.Next() {
		var id int
		var name string
		var email string

		err = rows.Scan(&id, &name, &email)

		if err != nil {
			println(fmt.Sprintf("Error: %s", err))
			os.Exit(1)
		}

		println(fmt.Sprintf("ID: %d, Name: %s, Email: %s", id, name, email))
	}
}
