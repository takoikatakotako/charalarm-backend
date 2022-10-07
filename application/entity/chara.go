package entity

type Chara struct {
	CharaID          string       `json:"charaID" dynamodbav:"charaID"`
	CharaEnable      bool       `json:"charaEnable" dynamodbav:"charaEnable"`
	CharaName        string       `json:"charaName" dynamodbav:"charaName"`
	CharaDescription string       `json:"charaDescription" dynamodbav:"charaDescription"`
	CharaProfiles    []CharaProfile `json:"charaProfile" dynamodbav:"charaProfile"`
}

type CharaProfile struct {
	Title string `json:"title" dynamodbav:"title"`
	Name  string `json:"name" dynamodbav:"name"`
	URL   string `json:"url" dynamodbav:"url"`
}
