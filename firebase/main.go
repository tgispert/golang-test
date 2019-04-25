package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Brand ...
type Brand struct {
	ID       string `json:"id,omitempty"`
	Brand    string `json:"brand,omitempty"`
	Products []Product
}

// Product ...
type Product struct {
	Class  string `json:"class,omitempty"`
	Gender int    `json:"gender,omitempty"`
	Name   string `json:"name,omitempty"`
	Code   string `json:"code,omitempty"`
	Price  int    `json:"price,omitempty"`
}

func main() {
	opt := option.WithCredentialsFile("sdk/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	brand := Brand{ID: "2", Brand: "levis"}
	log.Print(brand)
	result, err := client.Collection("lista").Doc("brands").Set(context.Background(), brand)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(result)
	defer client.Close()
}
