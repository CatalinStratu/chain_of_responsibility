package main

import (
	"errors"
	"fmt"
	"time"
)

func test() (bool, error) {

	return false, errors.New("test")
}
func e(v int) (int, error) {
	if v == 0 {
		_, err := test()
		if err != nil {
			return 0, errors.New("Zero cannot be used+asdasdasdasd")
		}
		return 0, errors.New("Zero cannot be used")
	} else {
		return 2 * v, nil
	}
}

func main() {
	v, err := e(0)

	if err != nil {
		fmt.Println(err, v) // Zero cannot be used 0
	}

	t := time.Now()
	td := t.Format("2006-01-02 15:04:05")
	fmt.Println(td)
}

/*
type myWriter interface {
	Write(record []string) error
}

type myWriterImpl struct {
	err error
}

func (a myWriterImpl) Write(record []string) error {

	return a.err
}

*/
