package dataservice

import (
	"Project/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddResume(client *mongo.Client, w http.ResponseWriter, r *http.Request) error {
	var resume model.Resume
	if err := json.NewDecoder(r.Body).Decode(&resume); err != nil {
		return err
	}

	collection := client.Database("job_portal").Collection("resumes")

	_, err := collection.InsertOne(context.TODO(), bson.D{
		{Key: "id", Value: resume.UserID},
		{Key: "resume", Value: resume.Contents},
		{Key: "parsedop", Value: resume.ParsedOP},
	})
	if err != nil {
		return err
	}

	fmt.Println("User added successfully!")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resume)
	return nil
}
