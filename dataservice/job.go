package dataservice

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"Project/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddJob(db *mongo.Client, w http.ResponseWriter, r *http.Request) error {
	// Parse the request body to get the job data
	var job *model.JobListing
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return err
	}

	// Access the job collection
	collection := db.Database("job_portal").Collection("jobs")

	// Insert the job data into the database
	_, err = collection.InsertOne(context.Background(), job)
	if err != nil {
		http.Error(w, "Failed to add job to database", http.StatusInternalServerError)
		return fmt.Errorf("failed to add job to database: %w", err)
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	return nil
}

func GetJobByID(client *mongo.Client, jobID string) (model.JobListing, error) {
	// Access the job collection
	collection := client.Database("job_portal").Collection("jobs")

	// Define a filter for the job ID
	filter := bson.M{"job_id": jobID}

	// Perform the query to find the job by ID
	var job model.JobListing
	err := collection.FindOne(context.Background(), filter).Decode(&job)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.JobListing{}, fmt.Errorf("job not found")
		}
		return model.JobListing{}, fmt.Errorf("failed to find job: %w", err)
	}

	return job, nil
}
