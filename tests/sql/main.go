package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type product struct {
	id      int
	model   string
	company string
	price   int
}

func main() {
	db, err := sql.Open("mysql", "root:root@/gointernship")

	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	_, err = db.Exec("insert into gointernship.Products (model, company, price) values (?, ?, ?)",
		"iPhone X", "Apple", 72000)
	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	rows, err := db.Query("select * from gointernship.Products LIMIT 5,10")
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	var products []product

	for rows.Next() {
		p := product{}
		//err := rows.Scan(&p.id, &p.model, &p.company, &p.price)
		err := rows.Scan(&p.id, &p.model, &p.company, &p.price)

		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	for _, p := range products {
		fmt.Println(p.id, p.model, p.company, p.price)
	}
	//https://metanit.com/go/tutorial/10.2.php
}
