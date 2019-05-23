package main

import (
	"log"
	"net/http"
	"time"
	m "./middlewares"
	c "./controllers"

	h "github.com/Bebbolus/helpers"
)

//source server configuration struct to load from json configuration file
type server struct {
	Listento     string `json:"listento"`
	Readtimeout  int `json:"readtimeout"`
	Writetimeout int `json:"writetimeout"`
}

var ServerConf server

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
	h.ReadJson(&ServerConf, "configurations/server.json")

	//set server configurations
	srv := &http.Server{
		ReadTimeout:  time.Duration(ServerConf.Readtimeout) * time.Second,
		WriteTimeout: time.Duration(ServerConf.Writetimeout) * time.Second,
		Addr:         ServerConf.Listento,
	}

	// mchain = append(chain, m.Method("GET")) //change method to thest that it work!
	http.HandleFunc("/home", Chain(c.Home, []mid{m.Method("GET"), m.SessIdHtml(true)}))
	http.HandleFunc("/estrazione", Chain(c.Estrazione, []mid{m.Method("POST"), m.SessIdHtml(false)}))
	
	http.HandleFunc("/img/", c.Static)
	http.HandleFunc("/css/", c.Static)

	log.Println("start HTTP listening on ", ServerConf.Listento)
	//SERVER START AND ERROR MANAGEMENT
	//best practise: start a local istance of server mux to avoid imported lib to define malicious handler
	log.Fatal(srv.ListenAndServe(), http.NewServeMux())
}