package main

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	db, err := sql.Open("sqlserver", "sqlserver://sa:P@ssw0rd@13.76.163.73/?database=techcoach")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	query := "SELECT * FROM cover"
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	id := 0
	name := ""
	ok := rows.Next()
	if ok {
		rows.Scan(&id, &name)
	}

	println(id, name)
}
