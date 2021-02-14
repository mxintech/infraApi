package server

import (
	"fmt"
	"log"
	"net/http"
)

type handler struct{}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Stuff", r.RequestURI)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("si"))
}

func StartServer() {
	mux := http.NewServeMux()
	mux.Handle("/api/", handler{})

	log.Fatal(http.ListenAndServe(":3000", mux))
}
