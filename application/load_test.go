package application

import (
	"github.com/google/go-cmp/cmp"
	"log"
	"os"
	"reflect"
	"testing"
)

//Users
var u1 = user{id: "1", firstName: "Test First Name1", lastName: "Test last Name1", email: "test1@test.com", gender: "male", ipAddress: "69.178.172.124"}
var u2 = user{id: "2", firstName: "Test First Name2", lastName: "Test last Name2", email: "test2@test.com", gender: "male", ipAddress: "69.178.172.124"}
var u3 = user{id: "3", firstName: "Test First Name3", lastName: "Test last Name3", email: "test3@test.com", gender: "male", ipAddress: "69.178.172.124"}
var u4 = user{id: "4", firstName: "Test First Name4", lastName: "Test last Name4", email: "test4@test.com", gender: "male", ipAddress: "69.178.172.124"}
var u5 = user{id: "5", firstName: "Test First Nam5", lastName: "Test last Nam5", email: "test5@test.com", gender: "male", ipAddress: "69.178.172.124"}
var u6 = user{id: "6", firstName: "Test First Name6", lastName: "Test last Name6", email: "test6@test.com", gender: "male", ipAddress: "69.178.172.124"}
var u7 = user{id: "7", firstName: "Test First Name7", lastName: "Test last Name7", email: "test7@test.com", gender: "male", ipAddress: "69.178.172.124"}
var u8 = user{id: "8", firstName: "Test First Name8", lastName: "Test last Name8", email: "test8@test.com", gender: "male", ipAddress: "69.178.172.124"}
var u9 = user{id: "9", firstName: "Test First Name9", lastName: "Test last Name9", email: "test9@test.com", gender: "male", ipAddress: "69.178.172.124"}
var u10 = user{id: "10", firstName: "Test First Name10", lastName: "Test last Name10", email: "test10@test.com", gender: "male", ipAddress: "69.178.172.124"}

func TestChunkName(t *testing.T) {
	name := chunkName(1)

	if name != "./chunk_1.csv" {
		t.Errorf("Chunk name was incorrect, got: %v, want: %d.", name, 10)
	}
}

func TestTableHeaderLine(t *testing.T) {
	//Test 1
	header := []string{u1.id, u1.firstName, u1.lastName, u1.email, u1.gender, u1.ipAddress}

	// Test 2
	u2Expected := []string{u2.id, u2.firstName, u2.lastName, u2.email, u2.gender, u2.ipAddress}

	var tests = []struct {
		input    user
		expected []string
	}{
		{u1, header},
		{u2, u2Expected},
	}

	for _, test := range tests {
		if output := headerLine(test.input); !cmp.Equal(output, test.expected) {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func TestArrayChunk(t *testing.T) {
	var validLines []user
	validLines = append(validLines, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10)
	chunkSize := 2

	var divided [][]user

	chunks := arrayChunk(validLines, chunkSize, divided)
	var expected [][]user

	// Users slice
	usersSlice1 := []user{u1, u2}
	usersSlice2 := []user{u3, u4}
	usersSlice3 := []user{u5, u6}
	usersSlice4 := []user{u7, u8}
	usersSlice5 := []user{u9, u10}

	expected = append(expected, usersSlice1, usersSlice2, usersSlice3, usersSlice4, usersSlice5)

	if !reflect.DeepEqual(chunks, expected) {
		t.Errorf("Array chunk was incorrect, got: %v, want: %v.", expected, chunks)
	}
}

func TestWriteInChunkCheckErrors(t *testing.T) {
	var validLines []user
	validLines = append(validLines, u1, u2)
	//w := csv.NewWriter(os.Stdout)
	//w := myWriterImpl{err: errors.New("error")}
	//err := writeChunkInCSV(validLines, w)
	//if err != nil {
	//		return
	//	}

	// Write any buffered data to the underlying writer (standard output).
	//w.Flush()

	//	if err := w.Error(); err != nil {
	//		t.Errorf("Data could not be written")
	//	}
	//fmt.Errorf("asdasd")
}

func TestChunkGenerator(t *testing.T) {
	var divided [][]user

	// Users slice
	usersSlice1 := []user{u1, u2}

	header := []string{u1.id, u1.firstName, u1.lastName, u1.email, u1.gender, u1.ipAddress}

	divided = append(divided, usersSlice1)
	err := chunkGenerator(header, divided)
	if err != nil {
		return
	}

	if !fileExists("chunk_0.csv") {
		t.Errorf("Data could not be written")
	}

	e := os.Remove("chunk_0.csv")
	if e != nil {
		log.Fatal(e)
	}

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
