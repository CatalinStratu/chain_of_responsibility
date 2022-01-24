package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strings"
	"sync"
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

	var wg sync.WaitGroup

	value := func(chunk []string) {

		defer wg.Done()
		var values []interface{}
		queryStr := "INSERT INTO users (first_name, last_name, email, gender, ip_address) VALUES "

		for j := 0; j < len(chunk); j++ {
			st := strings.Split(chunk[j], "|")
			queryStr += "(?, ?, ?, ?, ?),"
			values = append(values, st[1], st[2], st[3], st[4], st[5])
		}

		queryStr = queryStr[0 : len(queryStr)-1]
		_, err = db.Exec(queryStr, values...)
		if err != nil {
			panic(err)
		}
	}

	size := 10000
	for i := 0; i < len(validLines); i += size {
		end := i + size
		if end > len(validLines) {
			end = len(validLines)
		}
		wg.Add(1)
		go value(validLines[i:end])
		wg.Wait()
	}

	var countRows int

	err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&countRows)
	switch {
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Number of rows are %s\n", count)
	}
}
