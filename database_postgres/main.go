package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "123456"
	dbname   = "dvdrental"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// sqlStatement := `
	// INSERT INTO users (age, email, first_name, last_name)
	// VALUES ($1, $2, $3, $4)
	// RETURNING id`
	// id := 0
	// err = db.QueryRow(sqlStatement, 20, "que@calhoun.io", "Ricardo", "Calhoun").Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("New record ID is:", id)
	sqlStatement := `
	UPDATE users
	SET first_name = $1, last_name = $2
	WHERE id = $3;`
	res, err := db.Exec(sqlStatement, "NewFirst", "NewLast", 1)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

}
