package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	cursor, err := mongoClient.Database("market").Collection("users").Find(context.Background(), bson.D{})
	if err != nil {
		return
	}

	var users []User
	if err := cursor.All(context.Background(), &users); err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Path[len("/api/getuser/"):]

	var user User
	query := bson.D{{"userid", userID}}
	if err := mongoClient.Database("market").Collection("users").FindOne(context.Background(), query).Decode(&user); err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return
	}

	created, err := mongoClient.Database("market").Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(created); err != nil {
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Path[len("/api/updateuser/"):]

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return
	}

	query := bson.D{{"userid", userID}}
	update := bson.D{{"$set", bson.D{{"username", user.Username}, {"email", user.Email}}}}
	updated, err := mongoClient.Database("market").Collection("users").UpdateOne(context.Background(), query, update)
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(updated); err != nil {
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Path[len("/api/deleteuser/"):]

	query := bson.D{{"userid", userID}}
	deleted, err := mongoClient.Database("market").Collection("users").DeleteOne(context.Background(), query)
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(deleted); err != nil {
		return
	}
}
