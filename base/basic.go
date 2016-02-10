package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"io/ioutil"
)
//hello world
//echo arguments
//find duplicate lines
func main() {
	fmt.Println("Ok,let's GO")


	echo()
	echo2()

	fmt.Println("argument list v3")
	fmt.Println(strings.Join(os.Args[1:], " "))
	//duplicatesFilesInput()
	duplicatesFilesOnly()
}


func echo() {
	var s, sep string
	fmt.Println("argument list")
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	fmt.Println("argument list v2")
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

//find duplicate lines in input

func duplicatesFilesInput() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	input.Err()
}

func duplicatesFilesOnly() {
	counts := make(map[string]int)
	for _, filename :=range os.Args[1:]{
	   data,err:= ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		for _,line:=range strings.Split(string(data),"\n"){
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
