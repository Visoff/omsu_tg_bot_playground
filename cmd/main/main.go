package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Visoff/omsu_eservice_api_fun/internal/router"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        panic(err)
    }
    ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
    defer cancel()

    token, ok := os.LookupEnv("TELEGRAM_TOKEN")
    if !ok {
        panic("TELEGRAM_TOKEN is not set")
    }
    b, err := bot.New(token)
    if err != nil {
        panic(err)
    }
    router.Init(b)
    b.Start(ctx)
}
