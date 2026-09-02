package ntfy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
	resp.Header.Set("Content-type", "application/json")

	res, err := c.client.Do(resp)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("server returned status %d", res.StatusCode)
	}

	return nil
}
