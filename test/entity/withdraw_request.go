package entity

type WithdrawRequest struct {
	UserID string `json:"userID"`
	UserToken string `json:"userToken"`
}
