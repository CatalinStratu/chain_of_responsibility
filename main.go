package main

import (
	"GoInternship_codeRefactoring/application"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {

	//inputs := &application.Inputs{FileName: "input.txt", ChunkSize: 1000, Type: "File"}
	inputs := &application.Inputs{FileName: "input.txt", ChunkSize: 100000, Type: "DataBase"}

	load := &application.Load{}
	extract := &application.Extract{}
	extract.SetNext(load)
	err := extract.Execute(inputs)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
