package main

import "fmt"

type A struct {
	data int
}

type Person struct {
	name string
	age int
	A
}

type Person2 struct {
	age int
}

type Employee struct {
	company string
	Person  //Anonymous member
	//we can access the anonymous member directly without reference of the person here
	Person2
}

//anonymous struct is strct without type keyword
//They are used when there is only a single use of the struct.
//It is often used with templates.
//It is often used with table driven tests.
var AnonymousStruct struct{
	config string
}


func main() {

	emp:=Employee{
		company: "test",
		Person:  Person{
			"testName",
			20,
			A{data:10}, // why we give , because in go compiler will add ; instead of us adding every statement in a line. so , means there is a continuation and so compiler will not add ;
		},
		Person2: Person2{age:25},
	}

	fmt.Println(emp.name)// we have got the value as employee.Name instead employee.person.name
	fmt.Println(emp.data)

	//fmt.Println(emp.age) // since we have Person and Person2 have the same field , compiler will throw an error saying that ambiguous field
	fmt.Println(emp.Person.age)

	//anonymous struct
	AnonymousStruct.config="TEST" //I cannot access any normal struct like this
	fmt.Println(AnonymousStruct)

	//second example of anonymous struct
	s := struct {
		name string
		age int
	}{
		"john",
		32,
	}
	fmt.Println(s)
}
