package main

import (
	"bufio"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	db, err := sql.Open("mysql", "root:root@/gointernship")

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtLines []string
	var validLines []string
	count := 0

	for scanner.Scan() {
		txtLines = append(txtLines, scanner.Text())
		if strings.Contains(txtLines[count], "||") == false {
			validLines = append(validLines, txtLines[count])
		}
		count++
	}

	value := func(chunk []string) {
		var values []interface{}
		queryStr := "INSERT INTO users (first_name, last_name, email, gender, ip_address) VALUES  "
		// slice de stringuri, nu de folosit Exec, de folosit ExecContext
		for j := 0; j < len(chunk); j++ {
			st := strings.Split(chunk[j], "|")
			queryStr += "(?, ?, ?, ?, ?),"
			values = append(values, st[1], st[2], st[3], st[4], st[5])
		}

		//queryStr = queryStr[0 : len(queryStr)-1]
		_, err = db.Exec(queryStr, values...)
		if err != nil {
			log.Fatal(err)
		}
	}

	size := 10000
	for i := 0; i < len(validLines); i += size {
		end := i + size
		if end > len(validLines) {
			end = len(validLines)
		}
		value(validLines[i:end])
	}
}
