package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strings"
)

var (
	ctx context.Context
)

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	db, err := sql.Open("mysql", "root:root@/gointernship")

	if err := db.Ping(); err != nil {
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				return
			}
		}(db)
		fmt.Println(fmt.Sprintf("database connection failed \"%v\"", err))
	}

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

	size := 10000
	for i := 0; i < len(validLines); i += size {
		end := i + size
		if end > len(validLines) {
			end = len(validLines)
		}
		insertChunk(validLines[i:end], db)
	}
}

// Insert chunk in Database
func insertChunk(chunk []string, db *sql.DB) {
	var values []interface{}
	queryStr := "INSERT INTO users (first_name, last_name, email, gender, ip_address) VALUES  "
	for j := 0; j < len(chunk); j++ {
		st := strings.Split(chunk[j], "|")
		queryStr += "(?, ?, ?, ?, ?),"
		values = append(values, st[1], st[2], st[3], st[4], st[5])
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	queryStr = queryStr[0 : len(queryStr)-1]
	_, err := db.ExecContext(ctx, queryStr, values...)
	if err != nil {
		log.Fatal(err)
	}
}
