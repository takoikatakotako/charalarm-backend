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
	DirectoryName string `json:"directoryName"`
	FileName      string `json:"fileName"`
}

type CharaExpression struct {
	Images []string `json:"images"`
	Voices []string `json:"voices"`
}

type CharaCall struct {
	Message string `json:"message"`
	Voice   string `json:"voice"`
}

//type CharaNameAndVoiceFileURL struct {
//	CharaName    string
//	VoiceFileName string
//}
