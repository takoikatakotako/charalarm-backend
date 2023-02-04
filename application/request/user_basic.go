package request

type UserSignUp struct {
	UserID    string `json:"userID"`
	UserToken string `json:"userToken"`
}