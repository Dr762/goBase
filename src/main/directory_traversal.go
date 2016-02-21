package main

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"fmt"
	"flag"
	"time"
	"sync"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")
var sema = make(chan struct{},20)//counting semafor for concurency release
var done = make(chan struct{})

func main() {

	//Determine inital directories
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//Traverse the file tree in parallel
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _,root :=range roots{
		n.Add(1)
		go walkDir(root,&n,fileSizes)
	}
	go func() {
                n.Wait()
		close(fileSizes)
	}()

	//print results
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64

	loop:
	for {
		select {

		case <-done:
		//Drain fileSizes to allow exit goroutine to finish
			for range fileSizes{

			}
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size

		case <-tick:
		//verbose
			printDiskUsage(nfiles, nbytes)


		}

	}

	//totals
	printDiskUsage(nfiles, nbytes)

	//cancel traversal when input is detected
	go func() {
		os.Stdin.Read(make([]byte,1)) //read a single byte
		close(done)
	}()
}

func walkDir(dir string, n *sync.WaitGroup,fileSizes chan <-int64) {
	defer n.Done()
	if cancelled(){
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n,fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	select {
	case sema <-struct {}{}:
	case <-done:
		return nil //cancelled

}
	defer func() {<-sema}()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "directory traversal: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes) / 1e9)

}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}