package response

type Chara struct {
	CharaID     string                     `json:"charaID"`
	Enable      bool                       `json:"enable"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Profiles    []CharaProfile             `json:"profiles"`
	Resources   []CharaResource            `json:"resources"`
	Expression  map[string]CharaExpression `json:"expressions"`
	Calls       []CharaCall                `json:"calls"`
}

type CharaProfile struct {
	Title string `json:"title"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type CharaResource struct {
	DirectoryName string `dynamodbav:"directoryName"`
	FileName      string `dynamodbav:"fileName"`
}

type CharaExpression struct {
	Images []string `dynamodbav:"images"`
	Voices []string `dynamodbav:"voices"`
}

type CharaCall struct {
	Message string `dynamodbav:"message"`
	Voice   string `dynamodbav:"voice"`
}

//type CharaNameAndVoiceFileURL struct {
//	CharaName    string
//	VoiceFileURL string
//}
