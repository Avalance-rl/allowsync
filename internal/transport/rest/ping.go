package rest

import (
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405)
		w.Write([]byte("GET-Method forbidden!"))
		return
	}
	w.Write([]byte("pong"))
}
