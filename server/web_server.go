package server

import (
	"fmt"
	"github.com/abondar24/GoBase/basic"
	"github.com/abondar24/GoBase/lissajous"
	"log"
	"net/http"
	"sync"
)

var count int
var mu sync.Mutex

func WebServer() {

	http.HandleFunc("/", simpleHandler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) { lissajous.Lissajous(w) })
	http.HandleFunc("/fractal", func(w http.ResponseWriter, r *http.Request) { basic.DrawFractal(w) })
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

//handler func
func simpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s/n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] =%q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
