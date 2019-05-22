package main

import (
	"log"
	"net/http"
	m "./middlewares"
	c "./controllers"
)

// middleware filter incoming HTTP requests.
// if the request pass the filter, it calls the next HTTP handler.
type mid func(http.HandlerFunc) http.HandlerFunc

// Chain provides mechanism to concatenate middlewares
func Chain(f http.HandlerFunc, mids []mid) http.HandlerFunc {
	for _, m := range mids {
		f = m(f)
	}
	return f
}

func main() {
	// mchain = append(chain, m.Method("GET")) //change method to thest that it work!
	http.HandleFunc("/", Chain(c.Home, []mid{m.Method("GET")}))
	log.Fatal(http.ListenAndServe(":8666", nil))
}