package main

import (
	"fmt"
	"time"
)

var imageCh =make(chan string, 10)
var videoCh =make(chan string, 10)
var newsCh =make(chan string, 10)

// we cant go write on closed channel
//once on a closed channel , read goes on continuing with default value till we stop with ok value in for loop kind of env -> so we should have ok to check before processing in select and all
//we can make a channel as nil for close, select cannot read from nil channel


func main(){

	go getImage(imageCh)
	go getVideo(videoCh)
	go getNews(newsCh)

	for {
		select{
		//here it goes on reading and the value wil be default after channel closed
			//so we have to check whether channel is closed or not
		case data, channelOpen:=<- newsCh:
			if channelOpen{
				fmt.Println(data)

			}else {
				newsCh=nil
			}

		case data,channelOpen:=<- videoCh:
			if channelOpen{
				fmt.Println(data)

			}else {
				videoCh=nil
			}
		case data,channelOpen:=<- imageCh:
			if channelOpen{
				fmt.Println(data)
			}else {
				imageCh=nil
			}
		case <-time.After(time.Minute):
			fmt.Println("Gateway Time out error ")


		}
		// anyhow if all the channles are closed, not use in going in this for loop
		// we cannot write on closed channels, so exiting that
		//one more thing is , in nil channel select case dont executed
		if videoCh == nil && newsCh == nil && imageCh == nil{
			fmt.Println("All channels are closed")
			break
		}
	}

	fmt.Println("The process is completed")
	
}

func getNews(ch chan string) {
	for i:=0;i<20;i++{
		ch<-"News"
	}
	close(ch)
}

func getVideo(ch chan string) {
	for i:=0;i<20;i++{
		ch<-"Video"
	}
	close(ch)
}

func getImage(ch chan string) {
	for i:=0;i<20;i++{
		ch<-"Image"
	}
	close(ch)
}
