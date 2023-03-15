package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Defining the customer struct
type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

// Hard coded Map of customers
var DATA = map[string]Customer{
	"61b80805-20bb-4a73-a8e0-ef3f413a9e6b": {
		ID:        "61b80805-20bb-4a73-a8e0-ef3f413a9e6b",
		Name:      "Pedro Gabriel",
		Role:      "Administrator",
		Email:     "pedro.gabriel@omano.ao",
		Phone:     "00244912325465",
		Contacted: false,
	},
	"800d5ea0-5914-495b-9138-b9548d38dfeb": {
		ID:        "800d5ea0-5914-495b-9138-b9548d38dfeb",
		Name:      "Rafael Nunes",
		Role:      "Sales Manager",
		Email:     "rafa.nunes@omano.ao",
		Phone:     "00244912325423",
		Contacted: true,
	},
	"ae331cad-6b03-4640-bebf-aaa43a1f3b8f": {
		ID:        "ae331cad-6b03-4640-bebf-aaa43a1f3b8f",
		Name:      "Irene Nunes",
		Role:      "Manager",
		Email:     "irene.nunes@omano.ao",
		Phone:     "00244934325423",
		Contacted: true,
	},
}

// Return a complete list of all customers
func getCustomers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if len(DATA) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(nil)
	} else {
		var customers []Customer
		for _, value := range DATA {
			customers = append(customers, value)
		}

		json.NewEncoder(w).Encode(customers)
		w.WriteHeader(http.StatusOK)
	}
}

// Return a single customer based on id
func getCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	customer, exits := DATA[id]

	if exits {
		json.NewEncoder(w).Encode(customer)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(nil)
	}
}

// Create a new customer
func addCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	id := uuid.New().String()

	if _, exist := DATA[id]; exist {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(nil)
	} else {
		var newEntry Customer
		reqBody, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(reqBody, &newEntry)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(nil)
		} else {
			newEntry.ID = id
			DATA[id] = newEntry
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newEntry)
		}
	}
}

// Update customer data
func updateCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if _, exits := DATA[id]; exits {
		var toUpdate Customer
		reqBody, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(reqBody, &toUpdate)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(nil)
		} else {
			toUpdate.ID = id
			DATA[id] = toUpdate
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(toUpdate)
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(nil)
	}
}

// Delete a specific customer based on id
func deleteCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if _, exits := DATA[id]; exits {
		delete(DATA, id)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(nil)
	}
}

//Serve a HTML page with endpoint informations
func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "<H3> CRM BankEnd API </H3>")
	http.ServeFile(w, r, "./static/home.html")
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")

	fmt.Println("Server running on port 7000 ...")
	log.Fatal(http.ListenAndServe(":7000", router))
}
