package matchedResume

import (
	"Project/model"
	"fmt"
	"sort"
	"strings"

	"Project/dataservice"

	"go.mongodb.org/mongo-driver/mongo"
)

func MatchResumeWithJob(db *mongo.Client, jobID string) ([]*model.MatchedResume, error) {
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
		return matchingResumes[i].MatchScore > matchingResumes[j].MatchScore
	})

	return matchingResumes, nil
}

func CalculateMatchScore(jobTechnicalSkills []string, resumeTechnicalSkills string) float64 {
	resumeSkills := extractSkillsFromString(resumeTechnicalSkills)
	matchingSkills := 0

	for _, jobSkill := range jobTechnicalSkills {
		if containsSkill(resumeSkills, strings.ToLower(jobSkill)) {
			matchingSkills++
		}
	}

	fmt.Println(matchingSkills)
	matchScore := (float64(matchingSkills) / float64(len(jobTechnicalSkills))) * 100
	return matchScore
}

// Improved skill extraction to handle various separators and normalize skill names
func extractSkillsFromString(technicalSkills string) []string {
	normalizedSkills := strings.ToLower(technicalSkills)
	normalizedSkills = strings.ReplaceAll(normalizedSkills, ", ", ",")
	normalizedSkills = strings.ReplaceAll(normalizedSkills, ",", ",")
	return strings.Split(normalizedSkills, ",")
}

// Adjusted to be case-insensitive and more flexible in matching
func containsSkill(skills []string, skill string) bool {
	skill = strings.ToLower(skill)
	for _, s := range skills {
		if strings.TrimSpace(s) == skill {
			return true
		}
	}
	return false
}
