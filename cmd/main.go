package main

import (
	"sdk/internal/app"
)

func main() {
	a := app.New()
	a.Start()
	defer a.Close()
}
