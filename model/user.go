package model

type User struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"Name,omitempty"`
	Email         string `json:"Email,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Authorization string `json:"authorization,omitempty"`
}
