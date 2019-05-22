package middlewares

import (
	"net/http"
	"strings"
)


func Method(args string) func(http.HandlerFunc) http.HandlerFunc {
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			//split args and check if the request as this method
			acceptedMethods := strings.Split(args, "|")
			for _, v := range acceptedMethods {
				if r.Method == v {
					// Call the next middleware in chain
					f(w, r)
					return
				}
			}
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}
}