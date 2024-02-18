package dataservice

import (
	"context"
	"fmt"

	"Project/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddResume(client *mongo.Client, resume *model.Resume) error {
	collection := client.Database("job_portal").Collection("resumes")
	resumeDoc := bson.D{
		{Key: "user_id", Value: resume.UserID},
		{Key: "technical_skills", Value: resume.TechnicalSkills},
	}
	_, err := collection.InsertOne(context.TODO(), resumeDoc)
	if err != nil {
		return fmt.Errorf("error adding resume to the database: %w", err)
	}

	fmt.Println("User added successfully!")
	return nil
}
