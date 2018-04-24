package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	var endpointList EndpointList
	endpointList.Endpoints = append(endpointList.Endpoints, "/generate/{numRecords}")

	respJSON, jsonErr := json.Marshal(endpointList)
	if jsonErr != nil {
		log.Errorf("failed to marshal json: %v", jsonErr)
		errorWithJSON(w, "an error occurred - unable to marshal endpointlist", http.StatusInternalServerError)
	}
	responseWithJSON(w, respJSON, http.StatusOK)
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	numRecords, convErr := strconv.Atoi(vars["numRecords"])
	if convErr != nil {
		log.Errorf("error generating records, could not convert input (%v) to integer: %v", vars["numRecords"], convErr)
		errorWithJSON(w, "an error occurred - unable to generate records", http.StatusInternalServerError)
	}

	recordSet, genErr := GenerateRecords(numRecords)
	if genErr != nil {
		log.Errorf("error generating records: %v", genErr)
		errorWithJSON(w, "an error occurred - unable to generate records", http.StatusInternalServerError)
	}

	respJSON, jsonErr := json.Marshal(recordSet)
	if jsonErr != nil {
		log.Errorf("failed to marshal json: %v", jsonErr)
		errorWithJSON(w, "an error occurred - unable to marshal generated records", http.StatusInternalServerError)
	}

	responseWithJSON(w, respJSON, http.StatusOK)

}

func errorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func responseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}
