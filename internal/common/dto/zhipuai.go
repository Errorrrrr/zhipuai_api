package dto

type AiRequest struct {
	Model       string    `json:"model"`
	Prompt      []*Prompt `json:"prompt"`
	TopP        float64   `json:"top_p"`
	Temperature float64   `json:"temperature"`
}

type AiResponse struct {
	Code    int64  `json:"code"`
	Msg     string `json:"msg"`
	Data    *Data  `json:"data"`
	Success bool   `json:"success"`
}

type Data struct {
	RequestId  string     `json:"request_id"`
	TaskId     string     `json:"task_id"`
	TaskStatus string     `json:"task_status"`
	Choices    []*Choices `json:"choices"`
	Usage      *Usage     `json:"usage"`
}

type Choices struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

type Prompt struct {
	RoleName string `json:"role_name"`
	Content  string `json:"content"`
}
