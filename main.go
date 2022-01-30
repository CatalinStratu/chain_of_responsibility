package main

import (
	"GoInternship_codeRefactoring/Application"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {

	//inputs := &Application.Inputs{FileName: "input.txt", ChunkSize: 1000, Type: "File"}
	inputs := &Application.Inputs{FileName: "input2.txt", ChunkSize: 2, Type: "DataBase"}

	load := &Application.Load{}
	extract := &Application.Extract{}

	extract.SetNext(load)
	err := extract.Execute(inputs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
