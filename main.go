package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	UID   string  `json:"UID"`
	Name  string  `json:"Name"`
	Desc  string  `json:"Desc"`
	Price float64 `json:"Price"`
}

//var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoints called: homepage")
}
func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function called: getInventory")
	json.NewEncoder(w).Encode(inventory)
}
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

}

type Listing struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Images      []string `json:"images"`
}

var inventory []Listing

func createListing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var listing Listing
	err := json.NewDecoder(r.Body).Decode(&listing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	inventory = append(inventory, listing)
	json.NewEncoder(w).Encode(listing)
}

func getListingByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, listing := range inventory {
		if listing.ID == id {
			json.NewEncoder(w).Encode(listing)
			return
		}
	}
	http.NotFound(w, r)
}

func getAllListings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventory)
}

func updateListingByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var listing Listing
	err = json.NewDecoder(r.Body).Decode(&listing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, item := range inventory {
		if item.ID == id {
			inventory[i] = listing
			json.NewEncoder(w).Encode(listing)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteListingByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, listing := range inventory {
		if listing.ID == id {
			inventory = append(inventory[:i], inventory[i+1:]...)
			json.NewEncoder(w).Encode(listing)
			return
		}
	}
	http.NotFound(w, r)
}

// func main() {
// 	inventory = append(inventory, Item{
// 		UID:   "0",
// 		Name:  "Cheese",
// 		Desc:  "A fine block of cheese",
// 		Price: 4.99,
// 	})
// 	inventory = append(inventory, Item{
// 		UID:   "0",
// 		Name:  "Milk",
// 		Desc:  "A jug of milk",
// 		Price: 3.25,
// 	})

// 	handleRequests()
// }
