package matchedResume

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetMatchingResumes(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var request struct {
			JobID string `json:"job_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		// Get job listing based on the provided job ID
		job, err := MatchResumeWithJob(db, request.JobID)
		if err != nil {
			http.Error(w, "Failed to retrieve job listing", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(job)
	}
}
