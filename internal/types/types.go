package types

type InvokeParams struct {
	Content string `json:"content"`
	UUID    string `json:"uuid"`
}

type DingMessageSendRequest struct {
	Text                      *text      `json:"text"`
	ConversationId            string     `json:"conversationId"`
	AtUsers                   []*atUsers `json:"atUsers"`
	ChatbotCorpId             string     `json:"chatbotCorpId"`
	ChatbotUserId             string     `json:"chatbotUserId"`
	MsgId                     string     `json:"msgId"`
	SenderNick                string     `json:"senderNick"`
	IsAdmin                   bool       `json:"isAdmin"`
	SenderStaffId             string     `json:"senderStaffId"`
	SessionWebhookExpiredTime int64      `json:"sessionWebhookExpiredTime"`
	CreateAt                  int64      `json:"createAt"`
	SenderCorpId              string     `json:"senderCorpId"`
	ConversationType          string     `json:"conversationType"`
	SenderId                  string     `json:"senderId"`
	ConversationTitle         string     `json:"conversationTitle"`
	IsInAtList                bool       `json:"isInAtList"`
	SessionWebhook            string     `json:"sessionWebhook"`
	RobotCode                 string     `json:"robotCode"`
	Msgtype                   string     `json:"msgtype"`
}

type text struct {
	Content string `json:"content"`
}

type atUsers struct {
	DingtalkId string `json:"dingtalkId"`
}
