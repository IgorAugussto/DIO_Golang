package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Order struct {
	ID        string   `json: "id, omitempty"`
	FirstName string   `json: "firstname, omitempty"`
	LastName  string   `json: "lastname, omitempty"`
	Email     string   `json: "email, omitempty"`
	Address   *Address `json: "address, omitempty"`
	Cakes     *Cakes   `json: cakes, omitempty`
}

type Address struct {
	City  string `json: "city, omitempty"`
	State string `json: "state, omitempty"`
}

type Cakes struct {
	Flavor string `json: "flavor, omitempty"`
	Size   string `json: "size, omitempty"`
}

var orders []Order

func GetOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(orders)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range orders {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Order{})
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	order.ID = params["id"]
	orders = append(orders, order)
	json.NewEncoder(w).Encode(orders)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range orders {
		if item.ID == params["id"] {
			orders = append(orders[:index], orders[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(orders)
	}
}

func main() {
	router := mux.NewRouter()

	orders = append(orders, Order{ID: "1", FirstName: "John", LastName: "Doe", Email: "jhon_doe@exemple.com",
		Address: &Address{City: "X", State: "X"}, Cakes: &Cakes{Flavor: "Chocolate", Size: "12"}})

	orders = append(orders, Order{ID: "2", FirstName: "Koko", LastName: "Doe", Email: "koko_doe@exemple.com",
		Address: &Address{City: "Y", State: "Y"}, Cakes: &Cakes{Flavor: "Orange", Size: "10"}})

	router.HandleFunc("/orders", GetOrders).Methods("GET")
	router.HandleFunc("/order{id}", GetOrder).Methods("GET")
	router.HandleFunc("/order/{id}", CreateOrder).Methods("POST")
	router.HandleFunc("/order/{id}", DeleteOrder).Methods("DELETE")
	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
