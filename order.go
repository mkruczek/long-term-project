package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type Order struct {
	OrderID       string  `json:"order_id"`
	Status        string  `json:"status"`
	Price         float64 `json:"price"`
	ProductID     string  `json:"product_id"`
	CustomerEmail string  `json:"customer_email"`
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	cursor, err := mongoClient.Database("market").Collection("order").Find(context.Background(), bson.D{})
	if err != nil {
		return
	}

	var products []Order
	if err := cursor.All(context.Background(), &products); err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		return
	}
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Path[len("/api/getorder/"):]

	var order Order
	query := bson.D{{"orderid", orderID}}
	if err := mongoClient.Database("market").Collection("order").FindOne(context.Background(), query).Decode(&order); err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(order); err != nil {
		return
	}
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		return
	}

	created, err := mongoClient.Database("market").Collection("order").InsertOne(context.Background(), order)
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(created); err != nil {
		return
	}
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		return
	}

	query := bson.D{{"orderid", order.OrderID}}
	update := bson.D{{"$set", bson.D{{"status", order.Status}}}}
	if _, err := mongoClient.Database("market").Collection("order").UpdateOne(context.Background(), query, update); err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(order); err != nil {
		return
	}
}
