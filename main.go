package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

var countCountries = []string{"United States of America", "United Kingdom"}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", handleIndex)
	router.HandleFunc("/generate/{numRecords}", handleGenerate).Methods("GET")

	log.Println("Listening on port 8080.")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("ERROR: Web server failed: error=(%v).\n", err)
	}
}
