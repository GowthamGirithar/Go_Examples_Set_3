package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type BookDetail struct {
	//exported field since it begins with a capital letter
	Name   string `json:"Name" xml:"Name"`
	Author string `json:"Author,omitempty"`
	CoAuthor string `json:"-"` // - field will not be serialized

}
//Name and name is same in go json Marshall,
func main() {

	//define an instance with required field
	b1 := BookDetail{
		Name:   "CS",
		Author: "Bakshi",
		CoAuthor: "None",

	}
	b2 := BookDetail{
		Name:   "The Red Sari",
		Author: "Javier",
		CoAuthor: "None",

	}

	books := make([]BookDetail, 0)
	books = append(books, b1, b2)

	bytes, err := json.Marshal(books)
	if err != nil {
		return
	}
	ioutil.WriteFile("test.txt", bytes, 0644)

	outputBooks := make([]BookDetail, 0)

	out, err := ioutil.ReadFile("test.txt")
	if err != nil {

		return

	}
	json.Unmarshal(out, &outputBooks)
	for _, bk := range outputBooks {
		fmt.Printf("the book is %+v", bk)
		fmt.Println()
	}

}
