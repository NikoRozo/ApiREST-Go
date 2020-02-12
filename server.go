package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nikorozo/apirest/datastore"
)

var (
	customers datastore.CustomerStore
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
	customers = &datastore.Customer{}
	customers.Initialize()
}

func main() {
	r := mux.NewRouter()
	log.Println("Customer api")
	api := r.PathPrefix("/customer").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "customer")
	})

	api.HandleFunc("/id/{id}", searchByCustomer).Methods(http.MethodGet)
	api.HandleFunc("/byScore", byScore).Methods(http.MethodGet)

	log.Fatalln(http.ListenAndServe(":8080", r))
}
