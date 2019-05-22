package controllers

import (
	"net/http"
	h "../helpers"
)

func Home(w http.ResponseWriter, r *http.Request) {
	sessionid := "stocazzo"
	myvar := map[string]interface{}{"session": sessionid}
	h.OutHtml(w, "assets/view/home.html", myvar)
}