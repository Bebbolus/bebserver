package controllers

import (
	"net/http"
	h "github.com/Bebbolus/helpers/web"
	hs "github.com/Bebbolus/helpers/sessions"
	"log"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("alphasession", r.FormValue("alphasession"))

	session := hs.Session{Key: r.Header.Get("key")}
	session.GetFromHeader(r)
	log.Println(session)
	h.OutLayout(w, "assets/view/home.html", nil)
}