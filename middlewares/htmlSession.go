package middlewares

import (
	"net/http"
	"net/url"
	h "github.com/Bebbolus/helpers"
)


func HtmlSession(init bool, data map[string]interface{}) func(http.HandlerFunc) http.HandlerFunc {
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			if r.FormValue("alphasession") == "" {
				if init{
					// data["sessid"] =  h.MakeUUID()

					v := url.Values{}
					v.Set("alphasession", h.MakeUUID())
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