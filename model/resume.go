package model

type Resume struct {
	UserID          string `bson:"user_id" json:"id"`
	TechnicalSkills string `bson:"technical_skills" json:"technicalSkills"`
}
