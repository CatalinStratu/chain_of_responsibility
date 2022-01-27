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

func (l *Load) Execute(i *Inputs) error {

	var divided [][]user
	divided = arrayChunk(users, i.ChunkSize, divided)

	err := chunkGenerator(headerLine(firstLine), divided)
	if err != nil {
		return errors.New("chunk generator error")
	}
	return nil
}

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

func (l *Load) SetNext(next step) {
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

	err = writeInChunk(divided[i], writer)
	if err != nil {
		return errors.New("write string error")
	}
	writer.Flush()
	return nil
}

func chunkGenerator(firstLine []string, divided [][]user) error {
	for i := 0; i < len(divided); i++ {
		err := createChunk(i, firstLine, divided)
		if err != nil {
			return err
		}
	}
	return nil
}

type myWriter interface {
	Write(record []string) error
}

type myWriterImpl struct {
	err error
}

func (a myWriterImpl) Write(record []string) error {

	return a.err
}

func writeInChunk(divided []user, writer myWriter) error {
	for j := 0; j < len(divided); j++ {
		u := []string{divided[j].id, divided[j].firstName, divided[j].lastName, divided[j].email, divided[j].gender, divided[j].ipAddress}
		err := writer.Write(u)
		if err != nil {
			return errors.New("write string error")
		}
	}
	return nil
}
