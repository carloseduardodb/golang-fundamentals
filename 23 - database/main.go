package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConnection := "root:RootPassword@/Backoffice?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected to database")

	rows, err := db.Query("select * from users;")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(id, name, email)
	}
}
