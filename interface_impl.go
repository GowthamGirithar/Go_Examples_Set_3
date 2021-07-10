package main

import "fmt"

type Test1 struct {
	data string
}
type Test2 struct {
	data string
}

type Tester interface {
	Print()
}

func (t Test1) Print(){
	fmt.Println("Hello")
}
//try list
func main() {
 t1:=Test1{}
 PrintData(t1)
 t2:=Test2{}
 PrintData(t2)

}

func PrintData(t interface{})  {
	//t should be of interface type
	val,ok:=t.(Tester)
	if ok{
		val.Print()
	}else{
		fmt.Println("Not implemented")
	}
}
