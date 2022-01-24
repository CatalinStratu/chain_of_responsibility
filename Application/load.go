package Application

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Load struct {
	next dates
}

func (l *Load) Execute(i *Inputs) {
	firstLineStr := []string{
		firstLine.id,
		firstLine.firstName,
		firstLine.lastName,
		firstLine.email,
		firstLine.gender,
		firstLine.ipAddress,
	}

	var divided [][]user

	divided = arrayChunk(users, i.ChunkSize, divided)

	writeInChunk(firstLineStr, divided)
	fmt.Println("Datele au fost impartite in chunk-uri")
}

func (l *Load) SetNext(next dates) {
	l.next = next
}

func arrayChunk(validLines []user, chunkSize int, divided [][]user) [][]user {
	for i := 0; i < len(validLines); i += chunkSize {
		end := i + chunkSize
		if end > len(validLines) {
			end = len(validLines)
		}
		divided = append(divided, validLines[i:end])
	}
	return divided
}

func chunkName(number int) string {
	return "./chunk_" + strconv.Itoa(number) + ".csv"
}

func writeInChunk(firstLine []string, divided [][]user) {
	for i := 0; i < len(divided); i++ {
		str := chunkName(i)
		f, e := os.Create(str)

		if e != nil {
			fmt.Println(e)
		}

		writer := csv.NewWriter(f)
		if e != nil {
			fmt.Println(e)
		}

		err := writer.Write(firstLine)
		if err != nil {
			log.Fatalf("Write string error %s", err)
		}

		for j := 0; j < len(divided[i]); j++ {
			u := []string{
				divided[i][j].id,
				divided[i][j].firstName,
				divided[i][j].lastName,
				divided[i][j].email,
				divided[i][j].gender,
				divided[i][j].ipAddress,
			}
			err := writer.Write(u)
			if err != nil {
				log.Fatalf("Write string error %s", err)
			}
		}

		writer.Flush()
	}
}
