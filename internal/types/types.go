package types

type InvokeParams struct {
	Content string `json:"content"`
	UUID    string `json:"uuid"`
}

type DingMessageSendRequest struct {
	Text           *text  `json:"text"`
	ConversationId string `json:"conversationId"`
}

type text struct {
	Content string `json:"content"`
}
