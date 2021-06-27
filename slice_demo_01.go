package main

import "fmt"

func main() {
	exampleOfArrayAndSlice()
	exampleOfSliceAndSlice()
	exampleOfCopyInSlice()
	exampleOfArrayCopy()

}

func exampleOfArrayCopy() {
	a1:=[...]int{1,2,3}
	//a2:=[...]int{1,2,3,4}
	//a2=a1 //compile time error only element of same length can be copied
	a3:=[...]int{10,20,30}
	a3=a1
	fmt.Println(a3)
}

func exampleOfCopyInSlice() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{}
	s3 := []int{1, 2, 3}
	s4 := []int{1, 2, 3, 4, 5, 6}
	s6 := make([]int, 5, 5)

	fmt.Println(s1, s2)
	fmt.Println(s2 == nil) // we have given []int{} , so it is not nil , if we have given as []int it will be nil
	//copy will copy the value and not the reference
	copy(s2, s1)
	//destination is empty , nothing is copied
	fmt.Println(s2)
	copy(s3, s1)
	//destination is less than src, length of des is copied
	fmt.Println(s3) //[1 2 3]
	copy(s4, s1)
	//destination is more in size than src, all are copied
	fmt.Println(s4) //[1 2 3 4 5 6]
        s4[0]=1000
	
	fmt.Println(s4,s1) //[1000 2 3 4 5 6] [1 2 3 4 5]
	//with make example, it have new memory and we are copying the value
	//but here it dint copy the reference
	copy(s6, s1)
	fmt.Println(s6)
	s6[0] = 100
	fmt.Println(s6, s1) //[100 2 3 4 5] [10 2 3 4 5]

}

func exampleOfSliceAndSlice() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[0:4] // s2 points to s1
	s3 := make([]int, 5, 5)

	fmt.Println(s1, s2, len(s2), cap(s2)) //[1 2 3 4 5] [1 2 3 4] 4 5

	//now change the value of s2[0], will change the value of s1[0]
	s2[0] = 10
	fmt.Println(s1, s2, len(s2), cap(s2)) //[10 2 3 4 5] [10 2 3 4] 4 5

	s2 = append(s2, 10, 20)
	fmt.Println(s1, s2, len(s2), cap(s2)) //[10 2 3 4 5] [10 2 3 4 10 20] 6 10

	//now s2 should have created a different memory
	s2[0] = 1
	fmt.Println(s1, s2, len(s2), cap(s2)) //[10 2 3 4 5] [1 2 3 4 10 20] 6 10

	// so slice always points to the one and on reaching the size it will create a new space

	//with make , it will  point to the memory location because we have created a memory with make ahead of assignment
	//slice will have the reference , so we are changing the reference
	s3 = s1[0:3]
	s3[0] = 100
	fmt.Println(s1, s3, len(s3), cap(s3)) //[100 2 3 4 5] [100 2 3] 3 5

}

func exampleOfArrayAndSlice() {
	//example of array and slice
	a := [5]int{1, 2, 3, 4, 5} // array -and other form of array declaration is [...]int{1,2,3,4,5} if we define slice it will be like a:=[]int{1,2,3,4,5}
	s := a[0:4]                // slice always points to array begining index of 0 and end index is 3 in this case , so len is 4

	fmt.Println(a, s, len(s), cap(s)) //[1 2 3 4 5] [1 2 3 4] 4 5
	//fmt.Println(s[4]) // it will throw index out of bound exception- becuase only length is size and cap is just for internal use so that it dont need to allocate the memory often

	//now change the value of s[0], will change the value of a[0] because slice is pointing to array only
	s[0] = 10
	fmt.Println(a, s, len(s), cap(s)) //[10 2 3 4 5] [10 2 3 4] 4 5

	// we have now length of s is 4, we will add one value
	s = append(s, 90) //[10 2 3 4 90] [10 2 3 4 90] 5 5
	fmt.Println(a, s, len(s), cap(s))
	// now we have reached the array size, if we add again it wil copy the value and create in new space,

	s = append(s, 90)
	fmt.Println(a, s, len(s), cap(s)) //[10 2 3 4 90] [10 2 3 4 90 90] 6 10

	//since as per our understanding , slice created a new copy and working in different addreess, so let change the value of s[0]
	//you will not see a value have changed
	s[0] = 1
	fmt.Println(a, s, len(s), cap(s)) //[10 2 3 4 90] [1 2 3 4 90 90] 6 10

	//so as per understanding, slice points to array in this case and then if it reaches the array size , it will create a new memeory and copy the contents
}
