package jobListing

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(router *mux.Router, db *mongo.Client) {
	router.HandleFunc("/createjob", createJob(db))
	router.HandleFunc("/getjob/{id}", getJob(db)).Methods("GET")
}
