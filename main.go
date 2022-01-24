package main

import (
	"GoInternship_codeRefactoring/Application"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	inputs := &Application.Inputs{FileName: "input.txt", ChunkSize: 1000, Type: "File"}
	//inputs := &Application.Inputs{FileName: "input.txt", ChunkSize: 10000, Type: "DataBase"}

	load := &Application.Load{}
	extract := &Application.Extract{}
	extract.SetNext(load)
	extract.Execute(inputs)
}
