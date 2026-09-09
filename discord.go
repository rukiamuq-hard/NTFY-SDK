package ntfy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Discord(ctx context.Context, dr DiscordRequest) error {
	data, err := json.Marshal(dr)
	if err != nil {
		return err
	}

	resp, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/api/ds", bytes.NewReader(data))
	if err != nil {
		return err
	}
	resp.Header.Add("Content-type", "application/json")

	res, err := c.client.Do(resp)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("server returned status %d", res.StatusCode)
	}

	fmt.Println("Sent is successful!")
	return nil
}
