package router

import (
	"github.com/Visoff/omsu_eservice_api_fun/internal/handler"
	"github.com/go-telegram/bot"
)

func Init(b *bot.Bot) {
    b.RegisterHandler(bot.HandlerTypeMessageText, "/schedule", bot.MatchTypeExact, handler.HandleModuleSelection)
    b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "select_group:", bot.MatchTypePrefix, handler.HandleGroupSelection)
    b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "schedule:", bot.MatchTypePrefix, handler.HandleSchedule)
}
