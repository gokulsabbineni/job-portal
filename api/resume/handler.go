package resume

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func addResume(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := AddResumeLogic(db, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func getResume(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract userID from URL parameters
		vars := mux.Vars(r)
		userID := vars["id"]
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		resume, err := GetResumeLogic(db, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal resume to JSON and send in response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resume)
	}
}
