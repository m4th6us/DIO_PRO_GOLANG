package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name string `json:"name"`
	Idade int `json:"idade"`
	Endereco string `json:"endereco"`
	Products []Products `json:"products"`
}

type Products struct {
	ProductName string `json:"product_name"`
	ProductPrice string `json:"product_price"`
	ProductCode string `json:"product_code"`
}

var Customers []Customer

func HandleReadCustomers(w http.ResponseWriter, r *http.Request) {

	log.Print("Read Customers")
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Customers)

}

func HandleCreateCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var customer Customer

	err := json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("Create Customer: ", customer)
	Customers = append(Customers, customer)

	json.NewEncoder(w).Encode(Customers)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/Customers", HandleReadCustomers).Methods("GET")
	r.HandleFunc("/Customers", HandleCreateCustomer).Methods("POST")

	http.ListenAndServe(":8090", r)
}