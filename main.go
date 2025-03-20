package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GET request received"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getHandler).Methods(("GET"))

	http.ListenAndServe(":8080", r)
}
