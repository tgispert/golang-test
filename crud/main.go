package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Brand struct {
	Id       string `json:"id,omitempty"`
	Brand    string `json:"brand,omitempty"`
	Products []Product
}

type Product struct {
	Class  string
	Gender int
	Name   string
	Code   string
	Price  int
}

var brands []Brand
var prod []Product

func GetBrand(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range brands {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Brand{})
}

func GetBrands(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(brands)
}

func CreateBrand(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var brand Brand
	_ = json.NewDecoder(req.Body).Decode(&brand)
	brand.Id = params["id"]
	brands = append(brands, brand)
	json.NewEncoder(w).Encode(brands)
}

func DeleteBrand(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range brands {
		if item.Id == params["id"] {
			brands = append(brands[:index], brands[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(brands)
}

func GetProduct(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/brand", GetBrands).Methods("GET")
	router.HandleFunc("/brand/{id}", GetBrand).Methods("GET")
	router.HandleFunc("/brand", CreateBrand).Methods("PUT")
	router.HandleFunc("/brand/{id}", DeleteBrand).Methods("DELETE")
	prod = append(prod, Product{Class: "", Gender: 0, Name: "", Code: "", Price: 0})
	brands = append(brands, Brand{Id: "1", Brand: "wrangler", Products: prod})
	log.Fatal(http.ListenAndServe(":12345", router))
}
