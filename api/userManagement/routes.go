package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(router *mux.Router, db *mongo.Client) {

	router.HandleFunc("/users", createUser(db))
	router.HandleFunc("/users/list/{id}", getUser(db)).Methods(http.MethodGet)
	router.HandleFunc("/users/update/{id}", func(w http.ResponseWriter, r *http.Request) { updateUser(db, w, r) }).Methods(http.MethodPut)
	router.HandleFunc("/resumes", addResume(db))
	//router.HandleFunc("/resumes/list/{id}", getResume(db)).Methods(http.MethodGet)

}
