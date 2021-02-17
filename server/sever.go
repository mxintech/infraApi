package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/TheGolurk/infraApi/api"
)

type handler struct{}

// ServeHTTP Handler to api
// more info and learning resources at: https://benhoyt.com/writings/go-routing/
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI

	switch path {
	case "/api/user/register":
		err := api.CreateUser(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found 😞"))
	}

}

// StartServer starts new server mux
func StartServer() {
	// WARNING!
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// mux := http.NewServeMux()
	// had := mux.Handle("/api/", handler{})

	server := http.Server{
		Addr:              ":3000",
		ReadTimeout:       2 * time.Minute,
		WriteTimeout:      2 * time.Minute,
		ReadHeaderTimeout: 2 * time.Minute,
		Handler:           handler{},
	}

	// log.Fatal(http.ListenAndServe(":3000", mux))
	log.Fatalln(server.ListenAndServe())
}
