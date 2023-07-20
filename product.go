package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type Product struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	cursor, err := mongoClient.Database("market").Collection("products").Find(context.Background(), bson.D{})
	if err != nil {
		return
	}

	var products []Product
	if err := cursor.All(context.Background(), &products); err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		return
	}
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Path[len("/api/getproduct/"):]

	var product Product
	query := bson.D{{"productid", productID}}
	if err := mongoClient.Database("market").Collection("products").FindOne(context.Background(), query).Decode(&product); err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(product); err != nil {
		return
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return
	}

	created, err := mongoClient.Database("market").Collection("products").InsertOne(context.Background(), product)
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(created); err != nil {
		return
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Path[len("/api/updateproduct/"):]

	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return
	}

	query := bson.D{{"productid", productID}}
	updated, err := mongoClient.Database("market").Collection("products").UpdateOne(context.Background(), query, product)
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(updated); err != nil {
		return
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Path[len("/api/deleteproduct/"):]

	query := bson.D{{"productid", productID}}
	deleted, err := mongoClient.Database("market").Collection("products").DeleteOne(context.Background(), query)
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(deleted); err != nil {
		return
	}
}
