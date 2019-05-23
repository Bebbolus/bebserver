package controllers

import (
	"net/http"
	"strings"
)

/*
	This function is like classic 
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("assets/img"))))
	but it manage all path from url and prevent navigation
*/
func Static(w http.ResponseWriter, r *http.Request) {	
	values := strings.Split(r.URL.Path, "/")
	prefix := "/"+values[1]+"/"
	fspath := "./assets/"+values[1]

	// fs := http.StripPrefix(prefix, http.FileServer(http.Dir(fspath)))
	fs := http.StripPrefix(prefix, http.FileServer(http.Dir(fspath)))
	if strings.HasSuffix(r.URL.Path, "/") {
		http.NotFound(w, r)
		return
	}
	fs.ServeHTTP(w, r)
}