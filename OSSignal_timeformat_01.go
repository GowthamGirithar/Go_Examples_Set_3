package main

import (
  "fmt"
  "log"
  "os"
  "os/signal"
  "syscall"
  "time"
)




func main() {

  exit:=make(chan int)
  //golang time format works as below
  //we need to use 2006 for year
  t := time.Now()
  log.Println("Correct:", t.Format("2006-01-02 15:04:05"))
  log.Println("Weird:", t.Format("2000-01-02 15:04:05"))

  WaitForsignal(exit)

  <- exit

}

func WaitForsignal(exit chan int) {
  sigs := make(chan os.Signal, 1)
  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)
   go func() {
     for sig := range sigs{
       switch sig {
       case syscall.SIGINT:
         fmt.Print("Hello")
         exit <-1
         return
       case syscall.SIGUSR1:
         fmt.Print("User 1")// we can initiate any thing here
         exit <-1
         return
       }
     }

   }()
}

