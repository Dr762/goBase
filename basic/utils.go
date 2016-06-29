package basic

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"time"
)

func Echo() {
	s, sep := "", ""
	fmt.Println("argument list ")
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func CountSha() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

}

//isn't it a closure?
func AnonFuncDemo() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func BigSlowOperation() {
	defer trace("bigslowOperation")()

	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Print("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
