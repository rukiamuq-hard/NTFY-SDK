<div align="center">

# NTFY-SDK

**A lightweight Go SDK for sending notifications through your own NTFY server**

[![Go Reference](https://img.shields.io/badge/go-reference-blue?logo=go)](https://github.com/rukiamuq-hard/NTFY-SDK)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.21-00ADD8?logo=go)](https://github.com/rukiamuq-hard/NTFY-SDK/blob/main/go.mod)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Stars](https://img.shields.io/github/stars/rukiamuq-hard/NTFY-SDK?style=social)](https://github.com/rukiamuq-hard/NTFY-SDK/stargazers)

Simple, dependency-free, production-ready.

</div>

---

## Features

- **Minimal design** — no external dependencies, only the Go standard library
- **Telegram support out of the box** — send messages to Telegram chats through your NTFY server
- **Full `context.Context` support** — timeouts and cancellation built in
- **Easy to extend** — straightforward to add new notification channels (Discord, Slack, Email, etc.)
- **Single client for all channels** — one entry point via `ntfy.New()`

---

## Installation

```bash
go get github.com/rukiamuq-hard/NTFY-SDK
```

Requires Go 1.21 or later.

---

## Quick Start

Import the package as `ntfy`:

```go
import "github.com/rukiamuq-hard/NTFY-SDK"
```

Example of sending a Telegram notification:

```go
package main

import (
	"context"
	"log"
	"time"

	ntfy "github.com/rukiamuq-hard/NTFY-SDK"
)

func main() {
	client := ntfy.New("localhost:8080")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := client.Telegram(ctx, ntfy.TelegramRequest{
		Token:   "Token",   // your bot token
		ChatID:  123123123, // your chat id
		Message: "message", // message text
	})
	if err != nil {
		log.Fatal(err)
		return
	}
}
```

---

## Usage

### Creating a client

```go
client := ntfy.New("localhost:8080")
```

`ntfy.New` accepts the base address of your NTFY server and returns a ready-to-use `*ntfy.Client`.

### Sending a Telegram notification

```go
err := client.Telegram(ctx, ntfy.TelegramRequest{
	Token:   "your-bot-token",
	ChatID:  123123123,
	Message: "Deployment completed successfully",
})
if err != nil {
	log.Fatal(err)
}
```

`Client.Telegram` sends a `POST` request to `{baseURL}/api/tg` with a JSON-encoded `TelegramRequest` body, and returns an error if the request fails or the server responds with a status outside the `2xx` range.

### Request struct

```go
type TelegramRequest struct {
	Token   string `json:"token"`
	ChatID  int64  `json:"chat_id"`
	Message string `json:"message"`
}
```

| Field     | Type     | Description                     |
|-----------|----------|----------------------------------|
| `Token`   | `string` | Telegram bot token               |
| `ChatID`  | `int64`  | Telegram chat or user ID         |
| `Message` | `string` | Message text to send             |

### Working with context

All client methods accept a `context.Context`, allowing you to control timeouts and cancellation:

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()

if err := client.Telegram(ctx, req); err != nil {
	// handle error/timeout
}
```

---

## Project Structure

```
NTFY-SDK/
├── client.go     # Client struct and New() constructor
├── models.go     # Request models (TelegramRequest)
├── telegram.go   # Client.Telegram implementation
└── go.mod
```

---

## Contributing

Pull requests and issues are welcome. If you'd like to add a new notification channel (Discord, Slack, Email, Webhook, etc.), feel free to open a pull request.

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/discord`)
3. Commit your changes (`git commit -m 'Add Discord support'`)
4. Push the branch (`git push origin feature/discord`)
5. Open a Pull Request

---

## License

No license file is currently included in the repository. Confirm licensing terms with the author before using this project in production.
