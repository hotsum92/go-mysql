package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:13306)/<db-name>?parseTime=true")

	if err != nil {
		println(fmt.Sprintf("Error: %s", err))
		os.Exit(1)
	}

	defer db.Close()

	result, err := db.Exec("show tables")

	if err != nil {
		println(fmt.Sprintf("Error: %s", err))
		os.Exit(1)
	}

	rows, err := result.RowsAffected()

	if err != nil {
		println(fmt.Sprintf("Error: %s", err))
		os.Exit(1)
	}

	println(fmt.Sprintf("Rows affected: %d", rows))
}
