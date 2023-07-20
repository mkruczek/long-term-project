package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

var mongoClient *mongo.Client

func health(w http.ResponseWriter, req *http.Request) {

	if err := mongoClient.Ping(context.Background(), nil); err != nil {
		http.Error(w, "unable to ping mongodb", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "i am alive\n")
}

func main() {

	http.HandleFunc("/api/health", health)

	http.HandleFunc("/api/getproducts", GetProducts)
	http.HandleFunc("/api/getproduct/", GetProduct)
	http.HandleFunc("/api/createproduct", CreateProduct)
	http.HandleFunc("/api/updateproduct/", UpdateProduct)

	http.HandleFunc("/api/getusers", GetUsers)
	http.HandleFunc("/api/getuser/", GetUser)
	http.HandleFunc("/api/createuser", CreateUser)
	http.HandleFunc("/api/updateuser/", UpdateUser)

	fmt.Println("starting server on port 8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(fmt.Sprintf("unable to start server: %v", err))
	}
}

func init() {

	//connect to mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:secret@localhost:27017"))
	if err != nil {
		panic(fmt.Sprintf("unable to connect to mongodb: %v", err))
	}
}
