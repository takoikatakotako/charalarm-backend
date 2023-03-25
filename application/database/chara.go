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
	CharaResources   []CharaResource            `dynamodbav:"resources"`
	CharaExpressions map[string]CharaExpression `dynamodbav:"expressions"`
	CharaCalls       []CharaCall                `dynamodbav:"calls"`
}

type CharaProfile struct {
	Title string `dynamodbav:"title"`
	Name  string `dynamodbav:"name"`
	URL   string `dynamodbav:"url"`
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
	Voices  string `dynamodbav:"voice"`
}
