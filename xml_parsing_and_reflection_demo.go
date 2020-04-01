package main

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

//demo of xml parser with pointer to members and normal one
//reflection to find the overridden method
type Test struct {
	Name    string  `xml:"name,attr"  json:"name,omitempty"`
	Age     int     `xml:"age,attr"  json:"age,omitempty"`
	Address *string `xml:"add,attr"  json:"add,omitempty"`
}

type PrintData interface {
	printData()
}

func (test Test) printData() {
	data, err := xml.Marshal(Test{})
	// * members value will be nil and xml parser will not consider nil value
	//other members value will be default value and so it considered
	if err == nil {
		fmt.Println(string(data))
	}

}

func main() {

	fm := Test{}
	dataType := reflect.TypeOf(fm) // return type
	fmt.Println("The data type is ", dataType)
	obj := reflect.New(dataType) // return the value
	inter := obj.Interface()
	fmt.Println("The interface value is ", inter)
	val, ok := inter.(PrintData) //value , ok:= i.(Type )
	if ok {
		val.printData()
	}

}
