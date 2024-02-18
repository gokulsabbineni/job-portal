package matchedResume

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(router *mux.Router, db *mongo.Client) {
	router.HandleFunc("/matchingresumes", GetMatchingResumes(db))
}
