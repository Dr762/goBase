package basic

import (
	"fmt"
	"os"
	"runtime"
)

func DeferPanicDemo() {
	defer printStack()
	f1(4)
}

func f1(x int) {
	fmt.Printf("f1(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f1(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
