package application

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Load struct {
	next step
}

// SetNext Set the next step
func (l *Load) SetNext(next step) {
	l.next = next
}

// Execute Function to divide users in chunks and write information in CSV files
func (l *Load) Execute(i *Inputs) error {
	var divided [][]user
	divided = arrayChunk(users, i.ChunkSize, divided)

	err := chunkGenerator(headerLine(firstLine), divided)
	if err != nil {
		return errors.New(fmt.Sprintf("chunk generator error"))
	}
	return nil
}

//Creates the CSV file header
func headerLine(line user) []string {
	header := []string{
		line.id,
		line.firstName,
		line.lastName,
		line.email,
		line.gender,
		line.ipAddress,
	}
	return header
}

//Based on how many lines should be in a chunk, it creates a chunk slice.
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

//Create the CSV file name
func chunkName(number int) string {
	return "./chunk_" + strconv.Itoa(number) + ".csv"
}

//Create the chunk
func createChunk(i int, firstLine []string, divided [][]user) error {
	str := chunkName(i)
	f, err := os.Create(str)
	if err != nil {
		return errors.New(fmt.Sprintf("error creating file: \"%v\"", err))
	}

	defer func(f *os.File) error {
		err := f.Close()
		if err != nil {
			return errors.New(fmt.Sprintf("error closing file: \"%v\"", err))
		}
		return nil
	}(f)

	writer := csv.NewWriter(f)
	if err != nil {
		return errors.New(fmt.Sprintf("error creating new writer: \"%v\"", err))
	}

	err = writer.Write(firstLine)
	if err != nil {
		return errors.New(fmt.Sprintf("write first line error: \"%v\"", err))
	}

	err = writeChunkInCSV(divided[i], writer)
	if err != nil {
		return errors.New(fmt.Sprintf("write chunk in CSV error: \"%v\"", err))
	}
	return nil
}

//Generate the chunk name
func chunkGenerator(firstLine []string, divided [][]user) error {
	for i := 0; i < len(divided); i++ {
		err := createChunk(i, firstLine, divided)
		if err != nil {
			return errors.New(fmt.Sprintf("Create chunk error: \"%v\"", err))
		}
	}
	return nil
}

//Slice of uses in CSV file
func writeChunkInCSV(divided []user, writer *csv.Writer) error {
	for j := 0; j < len(divided); j++ {
		u := []string{divided[j].id, divided[j].firstName, divided[j].lastName, divided[j].email, divided[j].gender, divided[j].ipAddress}
		err := writer.Write(u)
		if err != nil {
			return errors.New(fmt.Sprintf("Write in chunk error: \"%v\"", err))
		}
	}
	writer.Flush()
	return nil
}
