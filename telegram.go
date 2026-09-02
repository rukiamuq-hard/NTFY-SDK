package ntfy

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

func (c *Client) Telegram(ctx context.Context, req TelegramRequest) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/api/tg", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
