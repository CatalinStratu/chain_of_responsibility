package application

import (
	"bufio"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Extract structure
type Extract struct {
	next step
}

//User structure
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

//All users extracted from the database
var users []user

//CSV header
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
	} else {
		err := extract.extractDatesFromTxtFile(input)
		if err != nil {
			return errors.New(fmt.Sprintf("Could not retrieve text from file:\"%v\"", err))
		}
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
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				return
			}
		}(db)
		return errors.New(fmt.Sprintf("database connection failed \"%v\"", err))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rows, err := db.QueryContext(ctx, "SELECT * FROM users ORDER BY id")

	if err != nil {
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				return
			}
		}(rows)
		return errors.New(fmt.Sprintf("Query error \"%v\"", err))
	}

	for rows.Next() {
		u, err := readRowFromDB(rows)
		if err != nil {
			return errors.New(fmt.Sprintf("Error read data row from DB and \"%v\"", err))
		}
		users = append(users, u)
	}

	firstLine = users[0]
	users = users[1:]
	return nil
}

// Read row from DataBase
func readRowFromDB(rows *sql.Rows) (user, error) {
	u := user{}
	err := rows.Scan(&u.id, &u.firstName, &u.lastName, &u.email, &u.gender, &u.ipAddress)
	if err != nil {
		return u, errors.New(fmt.Sprintf("Read date from DataBase"))
	}
	return u, nil
}

// Extract dates from txt file
func (extract *Extract) extractDatesFromTxtFile(i *Inputs) error {
	file, err := os.Open(i.FileName)

	if err != nil {
		return errors.New(fmt.Sprintf("failed opening file: %s", err))
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
	return nil
}
