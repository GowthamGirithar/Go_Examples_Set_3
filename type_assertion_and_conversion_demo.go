package main

import "fmt"

func main() {

	// implicit conversion is not there in go
	// we have to do explicit converstion
	//Var i int32=5
	//Var j int64
	//j=I  // error even though it was bigger memory
	//
	//Consider -1 is i which is 4 bits and j is 8 bits and then the value will be 9 , that is wrong
	//So no implicit casting, we have to do explicit

	//this is conversion, so go consider it as im sending data to one method and it will convert and send it , so it is written as int32(i) , instead (int32)i
	var i int64=5
	var j int32
	//j=i - compile time error
	j=int32(i)
	fmt.Println(j)

	//type assertion - casting

	var data interface{}="hello"
	// i.(<Type>)
	if d,ok:=data.(string) ; ok{
		println(d)
	}

	switch a:=data.(type){ // we can use .(type) only in switch
	case string:
		println(a) // a is value
	}




	
}
