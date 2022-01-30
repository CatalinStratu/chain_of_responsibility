package Application

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
)

type Load struct {
	next step
}

func (l *Load) SetNext(next step) {
	l.next = next
}

// Execute Function to divide users in chunks and write information in CSV files
func (l *Load) Execute(i *Inputs) error {
	var divided [][]user
	divided = arrayChunk(users, i.ChunkSize, divided)

	err := chunkGenerator(headerLine(firstLine), divided)
	if err != nil {
		return errors.New("chunk generator error")
	}
	return nil
}

// Creates the CSV file header
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

//Based on how many lines a chunk should have, he creates a slice of chunks
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

func createChunk(i int, firstLine []string, divided [][]user) error {
	str := chunkName(i)
	f, e := os.Create(str)
	if e != nil {
		return errors.New("error creating file")
	}

	defer func(f *os.File) error {
		err := f.Close()
		if err != nil {
			return errors.New("error closing file")
		}
		return nil
	}(f)

	writer := csv.NewWriter(f)
	if e != nil {
		return errors.New("error creating new writer")
	}

	err := writer.Write(firstLine)
	if err != nil {
		return errors.New("write first line error")
	}

	err = writeChunkInCSV(divided[i], writer)
	if err != nil {
		return errors.New("write string error")
	}
	return nil
}

//Generate the chunk name
func chunkGenerator(firstLine []string, divided [][]user) error {
	for i := 0; i < len(divided); i++ {
		err := createChunk(i, firstLine, divided)
		if err != nil {
			return errors.New("test")
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
			return errors.New("test")
		}
	}
	writer.Flush()
	return nil
}
