package jobListing

import (
	"Project/dataservice"
	"Project/model"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func AddJobLogic(db *mongo.Client, w http.ResponseWriter, r *http.Request) error {
	return dataservice.AddJob(db, w, r)
}

func GetJobByID(client *mongo.Client, jobID string) (model.JobListing, error) {
	job, err := dataservice.GetJobByID(client, jobID)
	if err != nil {
		return model.JobListing{}, err
	}
	if job.JobID == "" {
		return model.JobListing{}, fmt.Errorf("job with id %s not found", jobID)
	}
	return job, nil
}
