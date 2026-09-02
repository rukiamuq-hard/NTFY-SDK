package ntfy

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func (c *Client) Telegram(req TelegramRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/api/tg", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
