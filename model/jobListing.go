package model

type JobListing struct {
	JobID      string   `bson:"job_id" json:"id"`
	JobDetails string   `bson:"job_details" json:"jobDetails"`
	JobTools   []string `bson:"job_tools" json:"jobTools"`
}
