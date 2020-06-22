package main

import (
	"os"
	"text/template"
)

type Book struct {
	//exported field since it begins with a capital letter
	Name   string
	Author string
}

func main() {
	//create a new template with some name
	t := template.New("Books template")

	//parse some content and generate a template,
	//which is an internal representation
	t, _ = t.Parse(`The books are
{{range .}}
----------------------------------------
Name: {{.Name}}
Author:   {{.Author}}
{{end}}
	`)

	//define an instance with required field
	b1 := Book{
		Name:   "CS",
		Author: "Bakshi",
	}
	b2 := Book{
		Name:   "The Red Sari",
		Author: "Javier",
	}

	books := make([]Book, 0)
	books = append(books, b1, b2)

	t.Execute(os.Stdout, books)
}
