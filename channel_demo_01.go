package main

import (
	"sync"
	"time"
)

//when the buffer size is 0, read will be blocked
//when the buffer size is 0, write want to be connect to read kind of, otherwise deadlock
func main() {

	case10()

}

// read here throws deadlock error, because no one is there to write
func case1(){
	ch:=make(chan int, 0) //by default 0
	<-ch  // it will see no one is writing till now and so what is point in reading , so throw deadlock

}

// read will block till it gets data
func case2(){
	ch:=make(chan int, 0) //by default 0
	go func() {
		ch <- 10
	}()
	<-ch
}

// read will throw deadlock, because no one is there to write the data
func case3(){
	ch:=make(chan int, 0) //by default 0
	go func() {

	}()
	<-ch
}

// read will block till it gets data and success
func case4(){
	ch:=make(chan int, 0) //by default 0
	go testGo4(ch)
	<-ch
}

func testGo4(c chan int){
	c <- 10
}

//write blocks if there is channel full
func case5(){
	ch:=make(chan int, 0)
	//deadlock error- because it is more than buffer size
	ch <- 10
	ch <- 11
	<- ch
	<- ch

}

func case6(){
	ch:=make(chan int, 0)
	//deadlock error- because it is more than buffer size and this time no one to read
	ch <- 10
	<- ch
	ch <- 11

	<- ch

}
//success , because someone to read
func case7(){
	ch:=make(chan int, 0)
	go readData(ch)
	ch <- 10
}

func readData(ch chan int) {
<-ch
}

//go blocks on write if buffer is full
func case8(){
	ch:=make(chan int, 3)
	var wg sync.WaitGroup
	wg.Add(2)
	go readDatas(ch, &wg)
	go writeDatas(ch,&wg)
	wg.Wait()
}

func writeDatas(ch chan int, wg *sync.WaitGroup) {
	ch<- 19
	ch<-19
	ch<-19
	ch<-19
	wg.Done()
}

func readDatas(ch chan int,wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond)
	for i:=0 ; i<4;i++{
		println(<-ch)
	}
	/**
	//deadlock error, when it goes to ch read after 4th one, write go routine is not there , so deadlock
	for i:=range ch{
		print(i)
	}
	 */
	wg.Done()
}

//cant write on closed channel , but we can read on closed channel which wll be in select
func case9(){
	ch:=make(chan int, 3)
	close(ch)
	ch<-10
}
// we can read on the closed channel
func case10(){
	ch:=make(chan int, 3)
	ch<- 10
	close(ch)
	<-ch
}
