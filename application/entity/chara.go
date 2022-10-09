package entity

type Chara struct {
	CharaID          string                     `json:"charaID" dynamodbav:"charaID"`
	CharaEnable      bool                       `json:"charaEnable" dynamodbav:"charaEnable"`
	CharaName        string                     `json:"charaName" dynamodbav:"charaName"`
	CharaDescription string                     `json:"charaDescription" dynamodbav:"charaDescription"`
	CharaProfiles    []CharaProfile             `json:"charaProfile" dynamodbav:"charaProfile"`
	CharaResource    CharaResource              `json:"charaResource" dynamodbav:"charaResource"`
	CharaExpression  map[string]CharaExpression `json:"charaExpression" dynamodbav:"charaExpression"`
	CharaCall        CharaCall                  `json:"charaCall" dynamodbav:"charaCall"`
}

type CharaProfile struct {
	Title string `json:"title" dynamodbav:"title"`
	Name  string `json:"name" dynamodbav:"name"`
	URL   string `json:"url" dynamodbav:"url"`
}

type CharaResource struct {
	Images []string `json:"images" dynamodbav:"images"`
	Voices []string `json:"voices" dynamodbav:"voices"`
}

type CharaExpression struct {
	Images []string `json:"images" dynamodbav:"images"`
	Voices []string `json:"voices" dynamodbav:"voices"`
}

type CharaCall struct {
	Voices []string `json:"voices" dynamodbav:"voices"`
}
