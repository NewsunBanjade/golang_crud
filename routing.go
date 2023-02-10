package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRouting() {
	r := mux.NewRouter()

	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/employee", GetEmployee).Methods("GET")
	r.HandleFunc("/employee/{eid}", GetEmployeeById).Methods("GET")
	r.HandleFunc("/employee/{eid}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{eid}", DeleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
