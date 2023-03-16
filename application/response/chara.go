package response

type Chara struct {
	CharaID          string                     `json:"charaID"`
	CharaEnable      bool                       `json:"charaEnable"`
	CharaName        string                     `json:"charaName"`
	CharaDescription string                     `json:"charaDescription"`
	CharaProfiles    []CharaProfile             `json:"charaProfiles"`
	CharaResource    CharaResource              `json:"charaResource"`
	CharaExpression  map[string]CharaExpression `json:"charaExpression"`
	CharaCall        CharaCall                  `json:"charaCall"`
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
