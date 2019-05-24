package middlewares

import (
	"net/http"
	h "github.com/Bebbolus/helpers/sessions"
	"log"
)


func HeaderSession(init bool, key string) func(http.HandlerFunc) http.HandlerFunc {
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			session := h.Session{Key: key}
			
			if r.Header.Get("alphasession") == "" {
				if init{
					err := session.Init()
					if err != nil{
						log.Fatal(err)
						http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
						return
					}
					r.Header.Set("key",session.Key )
					r.Header.Set("alphasession",session.Crypt )
					
					// Call the next middleware in chain
					f(w, r)
					return
				} else {
					log.Fatal("no session!")
					http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
					return
				}
			} else {				
				f(w, r)
				return
			}		
		}
	}
}