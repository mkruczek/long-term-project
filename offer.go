package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type Offer struct {
	OfferID   string `json:"offer_id"`
	ProductID string `json:"product_id"`
	Price     string `json:"price"`
	UserID    string `json:"user_id"`
	Datetime  string `json:"datetime"`
}

// http CRUD methods for offers
func GetOffers(w http.ResponseWriter, r *http.Request) {
	cursor, err := mongoClient.Database("market").Collection("offers").Find(context.Background(), bson.D{})
	if err != nil {
		return
	}

	var offers []Offer
	if err := cursor.All(context.Background(), &offers); err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(offers); err != nil {
		return
	}
}

func GetOffer(w http.ResponseWriter, r *http.Request) {
	offerID := r.URL.Path[len("/api/getoffer/"):]

	var offer Offer
	query := bson.D{{"offerid", offerID}}
	if err := mongoClient.Database("market").Collection("offers").FindOne(context.Background(), query).Decode(&offer); err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(offer); err != nil {
		return
	}
}

func CreateOffer(w http.ResponseWriter, r *http.Request) {
	var offer Offer
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		return
	}

	created, err := mongoClient.Database("market").Collection("offers").InsertOne(context.Background(), offer)
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(created); err != nil {
		return
	}
}

func UpdateOffer(w http.ResponseWriter, r *http.Request) {
	var offer Offer
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		return
	}

	updated, err := mongoClient.Database("market").Collection("offers").UpdateOne(context.Background(), bson.D{{"offerid", offer.OfferID}}, bson.D{{"$set", offer}})
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(updated); err != nil {
		return
	}
}

func DeleteOffer(w http.ResponseWriter, r *http.Request) {
	offerID := r.URL.Path[len("/api/deleteoffer/"):]

	deleted, err := mongoClient.Database("market").Collection("offers").DeleteOne(context.Background(), bson.D{{"offerid", offerID}})
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(deleted); err != nil {
		return
	}
}
