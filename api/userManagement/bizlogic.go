package api

import (
	"Project/dataservice"
	"Project/model"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddLogic(db *mongo.Client, w http.ResponseWriter, r *http.Request) error {
	return dataservice.AddUser(db, w, r)
}

func GetUserByID(client *mongo.Client, userID int) (model.User, error) {
	users, err := dataservice.ListUser(client, bson.M{"id": userID})
	if err != nil {
		return model.User{}, err
	}
	if len(users) == 0 {
		return model.User{}, fmt.Errorf("User with ID %d not found", userID)
	}
	return users[0], nil
}

func UpdateLogic(db *mongo.Client, userID int, w http.ResponseWriter, r *http.Request) error {
	return dataservice.UpdateUser(db, userID, w, r)
}
