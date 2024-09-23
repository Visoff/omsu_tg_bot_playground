package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/Visoff/omsu_eservice_api_fun/pkg/eservice"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandleModuleSelection(ctx context.Context, b *bot.Bot, update *models.Update) {
    groups, err := eservice.Groups()
    if err != nil {
        b.SendMessage(ctx,
            &bot.SendMessageParams{
                ChatID: update.Message.Chat.ID,
                Text: "Что то пошло не так...попробуй ещё раз",
            },
        )
        return
    }
    markup := models.InlineKeyboardMarkup{
    	InlineKeyboard: [][]models.InlineKeyboardButton{},
    }
    modules := make(map[string]struct{})
    for _, group := range groups {
        module := strings.Split(group.Name, "-")[0]
        if _, ok := modules[module]; ok {
            continue
        }
        modules[module] = struct{}{}
        markup.InlineKeyboard = append(
            markup.InlineKeyboard,
            []models.InlineKeyboardButton{
                {
                    Text: module,
                    CallbackData: "select_group:"+module,
                },
            },
        )
    }

    _, err = b.SendMessage(ctx,
        &bot.SendMessageParams{
        	ChatID:               update.Message.Chat.ID,
        	Text:                 "Выберите модуль",
        	ReplyMarkup:          markup,
        },
    )
    if err != nil {
        fmt.Println("ERROR", err)
    }
}
