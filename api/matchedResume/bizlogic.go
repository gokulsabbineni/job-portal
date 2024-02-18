package matchedResume

import (
	"Project/model"
	"fmt"
	"sort"
	"strings"

	"Project/dataservice"

	"go.mongodb.org/mongo-driver/mongo"
)

func MatchResumeWithJob(db *mongo.Client, jobID string) ([]*model.Resume, error) {
	job, err := dataservice.GetJobByID(db, jobID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve job listing: %w", err)
	}

	resumes, err := dataservice.GetAllResumes(db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve resumes: %w", err)
	}

	// Extract the technical skills required for the job
	jobTechnicalSkills := job.JobTools
	var matchingResumes []*model.MatchedResume
	for _, resume := range resumes {
		resumeTechnicalSkills := resume.TechnicalSkills
		matchScore := CalculateMatchScore(jobTechnicalSkills, resumeTechnicalSkills)
		matchingResumes = append(matchingResumes, &model.MatchedResume{Resume: resume, MatchScore: matchScore})
	}

	sort.Slice(matchingResumes, func(i, j int) bool {
		return matchingResumes[i].MatchScore < matchingResumes[j].MatchScore
	})

	var sortedResumes []*model.Resume
	for _, matchedResume := range matchingResumes {
		sortedResumes = append(sortedResumes, matchedResume.Resume)
	}

	return sortedResumes, nil
}

func CalculateMatchScore(jobTechnicalSkills []string, resumeTechnicalSkills string) float64 {
	resumeSkills := extractSkillsFromString(resumeTechnicalSkills)
	matchingSkills := 0

	for _, jobSkill := range jobTechnicalSkills {
		if containsSkill(resumeSkills, jobSkill) {
			matchingSkills++
		}
	}

	matchScore := float64(matchingSkills) / float64(len(jobTechnicalSkills)) * 100
	return matchScore
}

// Helper function to split a string of technical skills into individual skills
func extractSkillsFromString(technicalSkills string) []string {
	return strings.Split(technicalSkills, ", ")
}

// Helper function to check if a slice of skills contains a particular skill
func containsSkill(skills []string, skill string) bool {
	for _, s := range skills {
		if s == skill {
			return true
		}
	}
	return false
}
