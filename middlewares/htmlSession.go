package middlewares

import (
	"net/http"
	"net/url"
	h "github.com/Bebbolus/helpers"
	"encoding/json"
	b64 "encoding/base64"
)

func HtmlSession(init bool, key string) func(http.HandlerFunc) http.HandlerFunc {
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			data := make(map[string]interface{})
			
			if r.FormValue("alphasession") == "" {
				if init{
					data["sessionid"] =  h.MakeUUID()
					marshalled, _ := json.Marshal(data)
					
					v := url.Values{}
					criptedVal := b64.StdEncoding.EncodeToString(h.Encrypt(marshalled, key))
					v.Set("alphasession",criptedVal )
					r.Form = v
					
					// Call the next middleware in chain
					f(w, r)
					return
				}
			} else {
				b64tobyte, _ := b64.StdEncoding.DecodeString(r.FormValue("alphasession"))
				b := h.Decrypt(b64tobyte, key)
				_ = json.Unmarshal(b, &data)

				v := url.Values{}
				for idx, val := range data{
					v.Set(idx, val.(string))
				}
				r.Form = v
				f(w, r)
				return
			}

			//return HTTP ERROR
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return			
		}
	}
}