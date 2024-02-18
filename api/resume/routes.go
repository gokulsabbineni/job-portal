package resume

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(router *mux.Router, db *mongo.Client) {
	router.HandleFunc("/resumes", addResume(db))
	//router.HandleFunc("/resumes/list/{id}", getResume(db)).Methods(http.MethodGet)
}
