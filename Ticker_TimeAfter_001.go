package main

import (
	"fmt"
	"time"
)


//difference  between ticker and time.after
func main() {
   ticker:=time.NewTicker(2*time.Second)
   for {
	   select {
   		case <- ticker.C:
   			fmt.Print("Hello")
   		case <- time.After( 6*time.Second):
				return
	}
   }

}

//output: hello hello hello hello hello...... goes one
//this means time after will check 6 seconds from last received one

