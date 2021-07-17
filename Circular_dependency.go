package main

import "fmt"

//circular dependency can be solved with this
//function call
func init(){
	caller=Te
}

var caller func(data string)string

func main_2(){
	fmt.Print(caller("hello"))
}

func Te(data string) string{
	return data
}
