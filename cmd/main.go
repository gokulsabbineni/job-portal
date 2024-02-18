package main

import (
	jobApi "Project/api/jobListing"
	matchApi "Project/api/matchedResume"
	resumeApi "Project/api/resume"
	userApi "Project/api/userManagement"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to MongoDB database
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		log.Println("here")
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Check the MongoDB connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	router := mux.NewRouter()
	userApi.RegisterRoutes(router, client)
	resumeApi.RegisterRoutes(router, client)
	jobApi.RegisterRoutes(router, client)
	matchApi.RegisterRoutes(router, client)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
