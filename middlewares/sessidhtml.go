package middlewares

import (
	"net/http"
	"net/url"
	h "github.com/Bebbolus/helpers/crypt"
)


func SessIdHtml(init bool) func(http.HandlerFunc) http.HandlerFunc {
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			if r.FormValue("sessid") == "" {
				if init{
					v := url.Values{}
					v.Set("sessid", h.MakeUUID())
    				r.Form = v
					// Call the next middleware in chain
					f(w, r)
					return
				}
			} else {
				f(w, r)
				return
			}

			//return HTTP ERROR
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return			
		}
	}
}