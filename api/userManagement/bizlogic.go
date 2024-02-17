package userManagement

import (
	"fmt"
	"job-portal/dataservice"
	"job-portal/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IBizlogic interface {
	AddLogic(user model.User) error
}

type Bizlogic struct {
	DB *mongo.Client
}

func NewBizlogic(db *mongo.Client) *Bizlogic {
	return &Bizlogic{DB: db}
}

func (bl *Bizlogic) AddLogic(user model.User) error {
	return dataservice.AddUser(bl.DB, user)
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
