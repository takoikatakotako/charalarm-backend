package entity

type SignUpRequest struct {
	UserID    string `json:"userID"`
	AuthToken string `json:"authToken"`
}
