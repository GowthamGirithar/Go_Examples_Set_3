package main

// we can define function as type
type f func(s string)string

// we can define function as variable
var f1 func(s int) int


func getData(s string)string{
	return s
}
//we can return function
func getDataIndirectly(s string)f{
	return func(s string) string {
		return "s"

	}
}


func main() {
	//here higher order function where the function name is passed
	perform(getData)
	//here getDataIndirectly have diff function syntax and so we have to invoke function which will return function and that match the type f
	perform(getDataIndirectly("test"))

	// we can use the varible f1 defined here
	//if you havent defined before then have to use like f1:=
	f1=func(s int) int {
		print(s)
		return s
	}
	f1(10)


}

//first class function where function can be passed as an arguement
func perform(f1 f){
	d :=f1("test")
	println(d)
}
