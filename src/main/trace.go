package main
import (
	"time"
	"log"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation()  {
   defer trace("bigslowOperation")()

	time.Sleep(10*time.Second)
}


func trace(msg string) func() {
	start :=time.Now()
	log.Print("enter %s",msg)
	return func(){
		log.Printf("exit %s (%s)",msg,time.Since(start))
	}
}