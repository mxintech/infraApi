package server

import (
	"crypto/tls"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/TheGolurk/infraApi/api"
	"github.com/TheGolurk/infraApi/db"
)

type handler struct {
	conn *sql.DB
}

// ServeHTTP Handler to api
// more info and learning resources at: https://benhoyt.com/writings/go-routing/
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI

	switch path {
	case "/api/user/register":
		err := api.CreateUser(w, r, h.conn)
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

	conn, err := db.GetDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}

	server := &http.Server{
		ReadTimeout:  5 * time.Millisecond,
		WriteTimeout: 10 * time.Millisecond,
		IdleTimeout:  120 * time.Millisecond,
		Handler:      &handler{conn: conn},
	}

	// log.Fatal(http.ListenAndServe(":3000", mux))
	log.Fatalln(server.ListenAndServe())
}
