package requests

type SendMessageRequest struct {
	Message string `json:"message"`
	From    string `json:"from"`
	To      string `json:"to"`
	Token   string `json:"token"`
}

type GetMessageRequest struct {
	MessageID string `json:"id"`
	To        string `json:"to"`
	Token     string `json:"token"`
}
