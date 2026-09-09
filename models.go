package ntfy

type TelegramRequest struct {
	Token   string `json:"token"`
	ChatID  int64  `json:"chat_id"`
	Message string `json:"message"`
}

type DiscordRequest struct {
	WHook   string `json:"webhook"`
	Message string `json:"message"`
}
