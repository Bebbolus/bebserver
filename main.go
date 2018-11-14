package main

import (
	"fmt"
	bootstrap "github.com/Bebbolus/gostron/bootstrap"
	configmanager "github.com/Bebbolus/gostron/libs"
	"log"
	"net/http"
	"os"
	"plugin"
	"strconv"
	"time"
)

//source configuration struct to map the json configuration file
type routes struct {
	Endpoints []struct {
		Handler     string `json:"handler"`
		Middlewares []struct {
			Handler string `json:"handler"`
			Params  string `json:"params"`
		} `json:"middlewares"`
		Path string `json:"path"`
	} `json:"endpoints"`
}

var RoutesConf routes

type server struct {
	Listento     string `json:"listento"`
	Readtimeout  string `json:"readtimeout"`
	Writetimeout string `json:"writetimeout"`
}

var ServerConf server

/* PLUGINS */

//local Http hanlder plugin interface
type Handler interface {
	Fire(w http.ResponseWriter, r *http.Request)
}

/* MIDDLEWARES */

//local Middlewares handler plugin interface
type Middleware interface {
	Pass()
}

func kill(msg interface{}) {
	fmt.Println(msg)
	os.Exit(1)
}

//start point
func main() {
	err := configmanager.ReadFromJson(&ServerConf, "configurations/server.json")
	if err != nil {
		kill(err)
	}

	err = configmanager.ReadFromJson(&RoutesConf, "configurations/routes.json")
	if err != nil {
		kill(err)
	}

	readtimeout, err := strconv.Atoi(ServerConf.Readtimeout)
	if err != nil {
		kill(err)
	}
	writetimeout, err := strconv.Atoi(ServerConf.Writetimeout)
	if err != nil {
		kill(err)
	}

	//SET UP SERVER TIMEOUT
	srv := &http.Server{
		ReadTimeout:  time.Duration(readtimeout) * time.Second,
		WriteTimeout: time.Duration(writetimeout) * time.Second,
		Addr:         ServerConf.Listento,
	}

	for _, v := range RoutesConf.Endpoints {
		// load module
		// 1. open the so file to load the symbols
		plug, err := plugin.Open(v.Handler)
		if err != nil {
			kill(err)
		}

		// 2. look up a symbol (an exported function or variable)
		// in this case, variable Controller
		symController, err := plug.Lookup("Handler")
		if err != nil {
			kill(err)
		}

		// 3. Assert that loaded symbol is of a desired type
		// in this case interface type Handler (defined above)
		var handler Handler
		handler, ok := symController.(Handler)
		if !ok {
			kill("unexpected type from module symbol")
		}

		var chain []bootstrap.Gate

		/*
		   per ogni middleware configurato da eseguire su questo path:
		       carica il plugin di quel middleware e lo carica
		       cerca la variabile per i parametri e gli assegna il valore come da configurazione
		       lo aggiunge alla catena
		*/

		for _, mid := range v.Middlewares {
			// load module
			// 1. open the so file to load the symbols
			plug, midErr := plugin.Open(mid.Handler)
			if midErr != nil {
				kill(midErr)
			}
			// 2. look up a symbol (an exported function or variable)
			// in this case, function Pass()
			symFunc, midErr := plug.Lookup("Pass")
			if midErr != nil {
				kill(midErr)
			}

			//chain = append(chain, symFunc.(func() bootstrap.Gate))
			chain = append(chain, symFunc.(func(string) bootstrap.Gate)(mid.Params))

		}
		// fine

		// 4. use the module to handle the request
		http.HandleFunc(v.Path, bootstrap.Chain(handler.Fire, chain...))

	}
	//best practise: start a local istance of server mux to avoid imported lib to define malicious handler
	mux := http.NewServeMux()

	//SERVER START AND ERROR MANAGEMENT
	log.Fatal(srv.ListenAndServe(), mux)

}
