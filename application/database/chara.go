package database

const (
	CharaTableName       = "chara-table"
	CHARA_TABLE_CHARA_ID = "charaID"
)

type Chara struct {
	CharaID          string                     `dynamodbav:"charaID"`
	CharaEnable      bool                       `dynamodbav:"enable"`
	CharaName        string                     `dynamodbav:"name"`
	CharaDescription string                     `dynamodbav:"description"`
	CharaProfiles    []CharaProfile             `dynamodbav:"profiles"`
	CharaResource    CharaResource              `dynamodbav:"resources"`
	CharaExpression  map[string]CharaExpression `dynamodbav:"expressions"`
	CharaCall        CharaCall                  `dynamodbav:"calls"`
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
