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

func GetResume(client *mongo.Client, userID string) (*model.Resume, error) {
	collection := client.Database("job_portal").Collection("resumes")

	var resume model.Resume
	filter := bson.M{"user_id": userID}

	err := collection.FindOne(context.Background(), filter).Decode(&resume)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Return nil if resume not found
		}
		return nil, fmt.Errorf("error retrieving resume: %w", err)
	}

	return &resume, nil
}

func GetAllResumes(client *mongo.Client) ([]*model.Resume, error) {
	collection := client.Database("job_portal").Collection("resumes")

	var resumes []*model.Resume

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error retrieving resumes: %w", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var resume model.Resume
		if err := cursor.Decode(&resume); err != nil {
			return nil, fmt.Errorf("error decoding resume: %w", err)
		}
		resumes = append(resumes, &resume)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error while retrieving resumes: %w", err)
	}

	return resumes, nil
}
