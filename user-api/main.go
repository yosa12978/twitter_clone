package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/yosa12978/twitter/user-api/app"
	"github.com/yosa12978/twitter/user-api/configs"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	configs.LoadConfig()
}

func main() {
	application := app.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := application.Run(ctx); err != nil {
		panic(err)
	}
}
