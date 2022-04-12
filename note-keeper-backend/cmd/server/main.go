package main

import (
	"note-keeper-backend/config"
)

func main() {
	cfg := config.Load()
	svc := registerService(cfg)

	svc.Run()
}
