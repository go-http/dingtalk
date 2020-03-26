package dingtalk

// 钉钉消息
type Message struct {
	MsgType string `json:"msgtype"`

	Markdown *MessagePartialMarkdown `json:"markdown,omitempty"`
	Text     *MessagePartialText     `json:"text,omitempty"`

	At *MessagePartialAt `json:"at,omitempty"`
}

// 钉钉消息AT部分
type MessagePartialAt struct {
	AtMobiles     []string `json:"atMobiles"`
	AtDingtalkIds []string `json:"atDingtalkIds"`
	IsAtAll       bool     `json:"isAtAll"`
}

// 钉钉消息文本部分
type MessagePartialText struct {
	Content string `json:"content"`
}

// 钉钉消息Markdown部分
type MessagePartialMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// 创建一条Markdown消息
func NewMarkdownMessage(title, text string) Message {
	partial := &MessagePartialMarkdown{
		Text:  text,
		Title: title,
	}
	msg := Message{
		MsgType:  "markdown",
		Markdown: partial,
	}

	return msg
}

// 创建一条文本消息
func NewTextMessage(content string) Message {
	partial := &MessagePartialText{
		Content: content,
	}
	msg := Message{
		MsgType: "text",
		Text:    partial,
	}

	return msg
}
