package main

import "fmt"

func main() {

	//i++
	//go doesnt support i++ in expression , they consider devloper dont want to use paper and pen to calculate something
	// it only supports i++ in seperate statement , so even ++i is considered as redudant

	i:=0
	i++
	println(i) //1
	//++i compile time error


	switch { // switch doesnt need to have expression here
		case i < 10:
			fmt.Println("inside i<10")
		case i < 5:
			fmt.Println("inside i<7")
		case i< 7:
			fmt.Println("inside i<7")


	}

	//output will be inside i<10

	//printing related

	fmt.Printf(" the data are %[2]v  and %[1]v" , 10, 20) // the data are 20  and 10 //%[]d value insdie [] is order of element is taken and starts from 1
	fmt.Printf(" the data are %*v" , 10, 20)//the data are         20   , 10 spaces here followed by value




}
