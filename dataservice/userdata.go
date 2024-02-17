package dataservice

import (
	"context"
	"encoding/json"
	"fmt"
	"job-portal/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddUser(client *mongo.Client, w http.ResponseWriter, r *http.Request) error {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err
	}

	collection := client.Database("job_portal").Collection("users")

	_, err := collection.InsertOne(context.TODO(), bson.D{
		{"id", user.ID},
		{"name", user.Name},
		{"email", user.Email},
		{"phone", user.Phone},
		{"authorization", user.Authorization},
	})
	if err != nil {
		return err
	}

	fmt.Println("User added successfully!")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	return nil
}

func ListUser(client *mongo.Client, filter bson.M) ([]model.User, error) {
	collection := client.Database("job_portal").Collection("users")
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var users []model.User
	for cursor.Next(context.TODO()) {
		var u model.User
		err := cursor.Decode(&u)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	fmt.Println("Users Listed successfully!")
	fmt.Println(users)
	return users, nil
}

func UpdateUser(client *mongo.Client, userID int, w http.ResponseWriter, r *http.Request) error {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err
	}

	collection := client.Database("job_portal").Collection("users")

	// Define the filter to identify the user by ID
	filter := bson.M{"id": userID}

	// Create a map to store update fields
	updateFields := make(map[string]interface{})

	// Check each field in the user struct
	if user.Name != "" {
		updateFields["name"] = user.Name
	}
	if user.Email != "" {
		updateFields["email"] = user.Email
	}
	if user.Phone != "" {
		updateFields["phone"] = user.Phone
	}
	if user.Authorization != "" {
		updateFields["authorization"] = user.Authorization
	}

	// Create the update document with $set
	update := bson.D{{"$set", updateFields}}

	// Perform the update operation
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("User with ID %d not found", userID)
	}

	fmt.Println("User updated successfully!")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	return nil
}
