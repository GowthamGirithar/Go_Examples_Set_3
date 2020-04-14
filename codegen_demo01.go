package main

import (
	"log"
	"os"
	"text/template"
)

var ProjectKey="TestProject"

func main() {

	var t *template.Template
	var f *os.File

	log.Print("Generate main()")

		t = LoadTemplate("maintest.gotmpl")
		createDir("cmd/" + "testCodeGen")

		f = createFile("cmd/testCodeGen/" + "testCodeGen" + ".go")
		defer func(f *os.File) {
			if err := f.Close(); err != nil {
				log.Panic("Failed to close file: ", err)
			}
		}(f)

	printdata := true

		err := t.Execute(f, struct {
			TemplateName          string
			ServiceName           string
			Printdata bool
		}{
			TemplateName:          t.Name(),
			ServiceName:           "testCodeGen",
			Printdata: printdata,
		})
		if err != nil {
			log.Panicf("Error processing template '%s'. Error: %s", t.Name(), err.Error())
		}
	}

func createFile(name string) *os.File {

	f, err := os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		log.Panicf("Failed to create file '%s' : %s", name, err.Error())
	}

	return f
}
func createDir(name string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		err := os.MkdirAll(name, 0777)
		if err != nil {
			log.Panicf("cannot make dir '%s' : %s", name, err)
		}
	}
}

// LoadTemplate loads and parses the specified golang template.
func LoadTemplate(name string) *template.Template {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Panic("Environment variable 'GOPATH' must be set.")
	}

	fname := gopath + "/src/" + ProjectKey + "/" + name
	log.Printf("Parsing template '%s'", fname)
	t, err := template.New(name).ParseFiles(fname)
	if err != nil {
		log.Panic("Invalid adgen template : ", err)
	}
	return t
}
