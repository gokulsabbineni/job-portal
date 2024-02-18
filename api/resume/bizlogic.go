package resume

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"

	"Project/dataservice"
	"Project/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func AddResumeLogic(db *mongo.Client, w http.ResponseWriter, r *http.Request) error {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing multipart form: "+err.Error(), http.StatusBadRequest)
		return err
	}

	file, _, err := r.FormFile("resume")
	if err != nil {
		http.Error(w, "Error retrieving the file: "+err.Error(), http.StatusBadRequest)
		return err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading uploaded file: "+err.Error(), http.StatusBadRequest)
		return err
	}

	cmd := exec.Command("python", "script\\pythonScript.py")
	var stdin bytes.Buffer
	stdin.Write(fileBytes)
	cmd.Stdin = &stdin

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("error running python script: %s, stderr: %s", err.Error(), stderr.String()), http.StatusInternalServerError)
		return err
	}
	technicalSkills := strings.TrimSpace(out.String())

	userID := r.FormValue("id")
	// encodedString := base64.StdEncoding.EncodeToString(fileBytes)
	resume := model.Resume{
		UserID: userID,
		// Contents:        encodedString,
		TechnicalSkills: technicalSkills,
	}
	err = dataservice.AddResume(db, &resume)
	if err != nil {
		http.Error(w, "Error adding resume to the database: "+err.Error(), http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resume)
	return nil
}

func GetResumeLogic(db *mongo.Client, userID string) (*model.Resume, error) {
	resume, err := dataservice.GetResume(db, userID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving resume from the database: %w", err)
	}

	if resume == nil {
		return nil, fmt.Errorf("resume not found for user ID: %s", userID)
	}

	return resume, nil
}
