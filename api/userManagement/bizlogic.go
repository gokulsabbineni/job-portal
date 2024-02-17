package api

import (
	"Project/dataservice"
	"Project/model"
	"fmt"
	"net/http"

	// "mime/multipart"

	"github.com/dslipak/pdf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddLogic(db *mongo.Client, w http.ResponseWriter, r *http.Request) error {
	return dataservice.AddUser(db, w, r)
}

func AddResumeLogic(db *mongo.Client, w http.ResponseWriter, r *http.Request) error {
	err := r.ParseMultipartForm(169922) // Adjust max size as needed
	if err != nil {
		http.Error(w, "error parsing multipart form data", http.StatusBadRequest)
	}
	file, _, err := r.FormFile("resume") // Change "resume" to your actual field name
	if err != nil {
		http.Error(w, "error retrieving uploaded file", http.StatusBadRequest)
	}
	defer file.Close()

	reader, err := pdf.NewReader(file, 169922)
	if err != nil {
		http.Error(w, "error reading file", http.StatusBadRequest)
	}
	text, err := reader.GetPlainText()
	if err != nil {
		http.Error(w, "error interpretting file", http.StatusBadRequest)
	}

	fmt.Println(text, "test")

	// // Example: Search for keywords like "skills", "experience", "education"
	// keywords := []string{"skills", "experience", "education"}
	// extractedInfo := make(map[string][]string)
	// for _, keyword := range keywords {
	// 	for _, line := range strings.Split(text, "\n") {
	// 		if strings.Contains(strings.ToLower(line), keyword) {
	// 			extractedInfo[keyword] = append(extractedInfo[keyword], line)
	// 		}
	// 	}
	// }
	return dataservice.AddResume(db, w, r)
}

func GetUserByID(client *mongo.Client, userID int) (model.User, error) {
	users, err := dataservice.ListUser(client, bson.M{"id": userID})
	if err != nil {
		return model.User{}, err
	}
	if len(users) == 0 {
		return model.User{}, fmt.Errorf("user with id %d not found", userID)
	}
	return users[0], nil
}

func UpdateLogic(db *mongo.Client, userID int, w http.ResponseWriter, r *http.Request) error {
	return dataservice.UpdateUser(db, userID, w, r)
}
