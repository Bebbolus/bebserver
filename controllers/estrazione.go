package controllers

import (
	"net/http"
	"fmt"
	// h "github.com/Bebbolus/helpers"
)

func Estrazione(w http.ResponseWriter, r *http.Request) {
	// myvar := map[string]interface{}{"session": r.FormValue("sessid")}
	fmt.Fprint(w, "SESSION CONTENT:", r.FormValue("sessid"))
}