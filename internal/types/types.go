package types

type InvokeParams struct {
	Content string `json:"content"`
	UUID    string `json:"uuid"`
}

type DingMessageSendRequest struct {
	Content string `json:"content"`
}
