package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	var wait sync.WaitGroup
	for i:=0;i<10;i++{
		wait.Add(1)   //change the number of total add to more than number of goroutines , then it will deadlock
		// deadlock will happen if you say wait for 10, but u start only 5 go routine
		go testPrint(wait, i)

	}

	wait.Wait()
	fmt.Println("operations is completed")
}

func testPrint(wait sync.WaitGroup, i int) {
	fmt.Println(i)
	time.Sleep(time.Millisecond)
	wait.Done()
	fmt.Println("operations is completed in done block for i" , i)
}
// if we dont pass the address, it will not decrement ,
//but it will throw panic since wait group is 0, it i and it will be 0, so it goes in negative
//panic: sync: negative WaitGroup counter - TODO


