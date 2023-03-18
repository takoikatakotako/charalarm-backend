package request

type AnonymousUserRequest struct {
	UserID    string `json:"userID"`
	AuthToken string `json:"authToken"`
}

// deprecated
