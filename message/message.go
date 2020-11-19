package message

const (
	MsgTypeText       = "text"
	MsgTypeLink       = "link"
	MsgTypeMarkdown   = "markdown"
	MsgTypeActionCard = "actionCard"
)

type TextMessage struct {
	MsgType string     `json:"msgtype"`
	Text    TextParams `json:"text"`
	At      AtParams   `json:"at"`
}

type TextParams struct {
	Content string `json:"content"`
}

type AtParams struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

type LinkMessage struct {
	MsgType string     `json:"msgtype"`
	Link    LinkParams `json:"link"`
}

type LinkParams struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageURL string `json:"messageUrl"`
	PicURL     string `json:"picUrl,omitempty"`
}

type MarkdownMessage struct {
	MsgType  string         `json:"msgtype"`
	Markdown MarkdownParams `json:"markdown"`
	At       AtParams       `json:"at"`
}

type MarkdownParams struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type ActionCardMessage struct {
	MsgType    string           `json:"msgtype"`
	ActionCard ActionCardParams `json:"actionCard"`
}

type ActionCardParams struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
	BtnOrientation string `json:"btnOrientation,omitempty"`
	HideAvatar     string `json:"hideAvatar,omitempty"`
}
