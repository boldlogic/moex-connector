package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/boldlogic/moex-connector/internal/application"
)

func main() {
	_, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	_, err := application.New()
	if err != nil {
		log.Fatalf("Не удалось создать приложение: %v", err)
	}

}
