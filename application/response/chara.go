package response

type Chara struct {
	CharaID     string                     `json:"charaID"`
	Enable      bool                       `json:"enable"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Profiles    []CharaProfile             `json:"profiles"`
	Resource    CharaResource              `json:"resources"`
	Expression  map[string]CharaExpression `json:"expressions"`
	CharaCall   CharaCall                  `json:"calls"`
}

type CharaProfile struct {
	Title string `json:"title"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type CharaResource struct {
	Images []string `json:"images"`
	Voices []string `json:"voices"`
}

type CharaExpression struct {
	Images []string `json:"images"`
	Voices []string `json:"voices"`
}

type CharaCall struct {
	Voices []string `json:"voices"`
}

type CharaNameAndVoiceFileURL struct {
	CharaName    string
	VoiceFileURL string
}
