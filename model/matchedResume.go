package model

type MatchedResume struct {
	Resume     *Resume `json:"resume"`
	MatchScore float64 `json:"match_score"`
}
