package database

type Chara struct {
	CharaID          string                     `dynamodbav:"ID"`
	CharaEnable      bool                       `dynamodbav:"charaEnable"`
	CharaName        string                     `dynamodbav:"charaName"`
	CharaDescription string                     `dynamodbav:"charaDescription"`
	CharaProfiles    []CharaProfile             `dynamodbav:"charaProfiles"`
	CharaResource    CharaResource              `dynamodbav:"charaResource"`
	CharaExpression  map[string]CharaExpression `dynamodbav:"charaExpression"`
	CharaCall        CharaCall                  `dynamodbav:"charaCall"`
}

type CharaProfile struct {
	Title string `dynamodbav:"title"`
	Name  string `dynamodbav:"name"`
	URL   string `dynamodbav:"url"`
}

type CharaResource struct {
	Images []string `dynamodbav:"images"`
	Voices []string `dynamodbav:"voices"`
}

type CharaExpression struct {
	Images []string `dynamodbav:"images"`
	Voices []string `dynamodbav:"voices"`
}

type CharaCall struct {
	Voices []string `dynamodbav:"voices"`
}

type CharaNameAndVoiceFileURL struct {
	CharaName    string
	VoiceFileURL string
}
