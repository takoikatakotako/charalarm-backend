package database

const (
	CharaTableName    = "chara-table"
	CharaTableCharaID = "charaID"
)

type Chara struct {
	CharaID          string                     `dynamodbav:"charaID"`
	Enable           bool                       `dynamodbav:"enable"`
	Name             string                     `dynamodbav:"name"`
	Description      string                     `dynamodbav:"description"`
	CharaProfiles    []CharaProfile             `dynamodbav:"profiles"`
	CharaResource    CharaResource              `dynamodbav:"resources"`
	CharaExpressions map[string]CharaExpression `dynamodbav:"expressions"`
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
