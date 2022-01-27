package Application

import (
	"bufio"
	"database/sql"
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

var users []user
var firstLine user

func (t *Extract) Execute(i *Inputs) {

	if i.extract {
		t.next.Execute(i)
		return
	}

	if i.Type == "DataBase" {
		t.dbExtract()
	} else if i.Type == "File" {
		t.txtExtract(i)
	} else {
		t.txtExtract(i)
	}

	i.extract = true
	t.next.Execute(i)
}

func (t *Extract) SetNext(next step) {
	t.next = next
}

func (t *Extract) dbExtract() {
	db, err := sql.Open("mysql", "root:root@/gointernship")

	err = db.Ping()
	if err != nil {
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("failed open DB: %s", err)
		}
	}(db)

	rows, err := db.Query("SELECT * FROM users ORDER BY id")

	if err != nil {
		log.Fatal(err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("failed load rows: %s", err)
		}
	}(rows)

	for rows.Next() {
		u := readUserFromDB(rows)
		users = append(users, u)
	}

	firstLine = users[0]
	users = users[1:]
}

func readUserFromDB(rows *sql.Rows) user {
	u := user{}
	err := rows.Scan(&u.id, &u.firstName, &u.lastName, &u.email, &u.gender, &u.ipAddress)
	if err != nil {
		fmt.Println(err)
	}
	return u
}
func (t *Extract) txtExtract(i *Inputs) {
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
