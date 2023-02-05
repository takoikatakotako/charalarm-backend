package entity

type SignUpRequest struct {
	UserID    string `json:"userID"`
	UserToken string `json:"userToken"`
}
