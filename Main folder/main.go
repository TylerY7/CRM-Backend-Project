package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        string
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

var customer1 = Customer{
	ID:        "01",
	Name:      "Tyler",
	Role:      "Prospective",
	Email:     "tzyoung@valdosta.edu",
	Phone:     "333-444-5555",
	Contacted: true,
}

var customer2 = Customer{
	ID:        "02",
	Name:      "Josh",
	Role:      "Premium",
	Email:     "tzyoung@valdosta.edu",
	Phone:     "111-222-3333",
	Contacted: false,
}
var customer3 = Customer{
	ID:        "03",
	Name:      "Easton",
	Role:      "Online",
	Email:     "tzyoung@valdosta.edu",
	Phone:     "777-888-9999",
	Contacted: false,
}

var customers = map[string]Customer{
	customer1.ID: customer1,
	customer2.ID: customer2,
	customer3.ID: customer3,
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	customer := customers[id]

	if _, exists := customers[id]; exists {
		json.NewEncoder(w).Encode(customer)

	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customers)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newCustomer Customer

	reqBody, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(reqBody, &newCustomer); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if _, exists := customers[newCustomer.ID]; exists {
		w.WriteHeader(http.StatusConflict)
		return
	}

	customers[newCustomer.ID] = newCustomer

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customers)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var updatedCustomer Customer

	reqBody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody, &updatedCustomer); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if _, exists := customers[id]; !exists {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	updatedCustomer.ID = id
	customers[id] = updatedCustomer

	json.NewEncoder(w).Encode(updatedCustomer)

}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	if _, exists := customers[id]; exists {
		delete(customers, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(customers)
	}

}

func main() {

	router := mux.NewRouter()

	//Displays index.html
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/static/", http.StripPrefix("/static/", fileServer))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	//Getting a single customer
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	//Getting all customers
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	//Creating a customer
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	//Updating a customer
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PATCH")
	//Deleting a customer
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Server is starting...")

	http.ListenAndServe(":3000", router)
}
