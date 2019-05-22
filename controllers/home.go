package controllers

import (
	"net/http"
	h "github.com/Bebbolus/helpers"
)

func Home(w http.ResponseWriter, r *http.Request) {
	myvar := map[string]interface{}{"sessid": r.FormValue("sessid")}
	h.OutHtml(w, "assets/view/home.html", myvar)
}