package types

type InvokeParams struct {
	Content string `json:"content"`
	UUID    string `json:"uuid"`
}

type DingMessageRequest struct {
	MsgParam           string `json:"msgParam"`
	MsgKey             string `json:"msgKey"`
	OpenConversationId string `json:"openConversationId"`
	RobotCode          string `json:"robotCode"`
	Token              string `json:"token"`
	CoolAppCode        string `json:"coolAppCode"`
}
