package response

type UserInfoResponse struct {
	UserID          string                  `json:"userID"`
	AuthToken       string                  `json:"authToken"`
	Platform        string                  `json:"platform"`
	IOSPlatformInfo IOSPlatformInfoResponse `json:"iOSPlatformInfo"`
}
