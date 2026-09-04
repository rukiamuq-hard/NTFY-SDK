<div align="center">

# NTFY-SDK

**A lightweight Go SDK for sending notifications through your own [NTFY-Service](https://github.com/rukiamuq-hard/NTFY-Service)**

[![Go Reference](https://img.shields.io/badge/go-reference-blue?logo=go)](https://github.com/rukiamuq-hard/NTFY-SDK)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.26.3-00ADD8?logo=go)](https://github.com/rukiamuq-hard/NTFY-SDK/blob/main/go.mod)
[![Stars](https://img.shields.io/github/stars/rukiamuq-hard/NTFY-SDK?style=social)](https://github.com/rukiamuq-hard/NTFY-SDK/stargazers)

![GitHub repo size](https://img.shields.io/github/repo-size/rukiamuq-hard/NTFY-SDK)
![GitHub last commit](https://img.shields.io/github/last-commit/rukiamuq-hard/NTFY-SDK)

Simple, dependency-free, production-ready.

</div>

## Features

- **Minimal design** — no external dependencies, only the Go standard library
- **Telegram support out of the box** — send messages to Telegram chats through your [NTFY-Service](https://github.com/rukiamuq-hard/NTFY-Service)
- **Full `context.Context` support** — timeouts and cancellation built in
- **Easy to extend** — straightforward to add new notification channels (Discord, Slack, Email, etc.)
- **Single client for all channels** — one entry point via `ntfy.New()`

## Installation

```bash
go get github.com/rukiamuq-hard/NTFY-SDK
```

Requires Go 1.26.3 or higher.

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
	}
}
```

## Usage

### Creating a client

```go
client := ntfy.New("localhost:8080")
```

`ntfy.New` accepts the base address of your [NTFY-Service](https://github.com/rukiamuq-hard/NTFY-Service) and returns a ready-to-use `*ntfy.Client`.

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
### Working with context

All client methods accept a `context.Context`, allowing you to control timeouts and cancellation:

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()

if err := client.Telegram(ctx, req); err != nil {
	// handle error/timeout
}
```

## Project Structure

```
NTFY-SDK/
├── client.go     # Client struct and New() constructor
├── models.go     # Request models (TelegramRequest)
├── telegram.go   # Client.Telegram implementation
└── go.mod
```
