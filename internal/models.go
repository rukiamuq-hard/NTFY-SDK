package ntfy

type TelegramRequest struct {
	Token   string `json:"token"`
	ChatID  int64  `json:"chat_id"`
	Message string `json:"message"`
}
