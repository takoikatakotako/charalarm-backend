package entity

type AnonymousUserRequest struct {
	UserID    string `json: "userID"`
	UserToken    string `json: "userToken"`
}
