package main

import (
	"fmt"
	"reflect"
)
//demo on runes
func main_3() {
	//reverse in batches
	arr := []rune{'A', 'B', 'C', 'D', 'E', 'F'}
	k := 3
	l := len(arr)

	for i := 0; i < l; i += k {
		for j := 0; j < k/2; j++ {
			temp := arr[i]
			arr[i] = arr[i+k-1]
			arr[i+k-1] = temp
		}
	}
	r:='A'
	fmt.Println(reflect.TypeOf(r)) //int 32
	// if we print the runes , it will print as int32 and to display as character we need to convert to string
	for _, ar := range arr {
		fmt.Println(string(ar)) // C B A F E D
		fmt.Println(ar) // 67 66 65 70 69 68
	}
}
