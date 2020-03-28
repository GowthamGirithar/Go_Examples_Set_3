package main

import "fmt"

type Emp struct{
	//reportsTo Emp  compile time error as recursive
	reportsTo *Emp  // pointer is allowed
	name string
}


func main() {
	var e1 Emp
	fmt.Println(e1)//nil


	e:=Emp{
		reportsTo: nil,
		name:      "Test",
	}

	fmt.Println("%#v" , e)
}


