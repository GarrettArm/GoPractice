// mminimal webserver that responds with your url as sample text

package main

import (
	"fmt"
	"log"
	"net/http"
)

var counts int

func main() {
	http.HandleFunc("/count/", counter)
	http.HandleFunc("/hello/", otherHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func counter(w http.ResponseWriter, r *http.Request) {
	counts += 1
	fmt.Fprintf(w, "counter = %v", counts)
}

func handler(w http.ResponseWriter, r *http.Request) {
	text := fmt.Sprintf("URL.Path = %q\nCount = %v", r.URL.Path, counts)
	fmt.Fprintf(w, text)
}

func otherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Other URL.Path = %q\n", r.URL.Path)
}
