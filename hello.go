package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	f := func() { fmt.Println("yo") }
	go f()
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
