package server

import (
	"log"
	"net/http"

	"github.com/TheGolurk/infraApi/api"
)

type handler struct{}

// ServeHTTP Handler to api
// more info and learning resources at: https://benhoyt.com/writings/go-routing/
func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI

	switch path {
	case "/api/user/register":
		api.CreateUser(w, r)

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found 😞"))
	}

}

// StartServer starts new server mux
func StartServer() {
	mux := http.NewServeMux()
	mux.Handle("/api/", handler{})

	log.Fatal(http.ListenAndServe(":3000", mux))
}