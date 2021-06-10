package app

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func AuthHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//r.URL.Host = url.Host
		//r.URL.Scheme = url.Scheme
		//r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		//r.Host = url.Host

		r.URL.Path = mux.Vars(r)["auth_path"]
		p.ServeHTTP(w, r)
	}
}
