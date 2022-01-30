package Application

import (
	"bufio"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Extract struct {
	next step
}

type user struct {
	id        string
	firstName string
	lastName  string
	email     string
	gender    string
	ipAddress string
}

var (
	ctx context.Context
)

var users []user
var firstLine user

func (extract *Extract) Execute(input *Inputs) error {
	if input.extract {
		err := extract.next.Execute(input)
		if err != nil {
			return errors.New(fmt.Sprintf("Could not extract data"))
		}
	}

	if input.Type == "DataBase" {
		err := extract.extractDatesFromDatabase()
		if err != nil {
			return errors.New(fmt.Sprintf("Could not retrieve text from database: \"%v\"", err))
		}
	} else if input.Type == "File" {
		extract.extractDatesFromTxtFile(input)
	} else {
		extract.extractDatesFromTxtFile(input)
	}

	input.extract = true
	err := extract.next.Execute(input)
	if err != nil {
		return errors.New(fmt.Sprintf("\"%v\"", err))
	}
	return nil
}

// SetNext Set the next step
func (extract *Extract) SetNext(next step) {
	extract.next = next
}

//Extract dates from Database
func (extract *Extract) extractDatesFromDatabase() error {
	db, err := sql.Open("mysql", "root:root@/gointernship")

	if err := db.Ping(); err != nil {
		defer db.Close()
		return errors.New("database connection failed")
	}

	rows, err := db.QueryContext(ctx, "SELECT * FROM users ORDER BY id")

	if err != nil {
		defer rows.Close()
		return errors.New("query error")
	}

	defer rows.Close()

	for rows.Next() {
		u := readUserFromDB(rows)
		users = append(users, u)
	}

	firstLine = users[0]
	users = users[1:]
	return nil
}

// Read row from DataBase
func readUserFromDB(rows *sql.Rows) user {
	u := user{}
	err := rows.Scan(&u.id, &u.firstName, &u.lastName, &u.email, &u.gender, &u.ipAddress)
	if err != nil {
		fmt.Println(err)
	}
	return u
}

// Extract dates from txt file
func (extract *Extract) extractDatesFromTxtFile(i *Inputs) {
	file, err := os.Open(i.FileName)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtLines []string
	count := 0

	for scanner.Scan() {
		txtLines = append(txtLines, scanner.Text())
		if strings.Contains(txtLines[count], "||") == false {
			line := strings.Split(txtLines[count], "|")
			u := user{id: line[0], firstName: line[1], lastName: line[2], email: line[3], gender: line[4], ipAddress: line[5]}
			users = append(users, u)
		}
		count++
	}

	firstLine = users[0]
	users = users[1:]
}
