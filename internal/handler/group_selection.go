package handler

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Visoff/omsu_eservice_api_fun/pkg/eservice"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandleGroupSelection(ctx context.Context, b *bot.Bot, update *models.Update) {
    module := strings.Split(update.CallbackQuery.Data, ":")[1]
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
    markup := models.InlineKeyboardMarkup{}
    today := time.Now().Format("02.01.2006")
    for _, group := range groups {
        if group.Name[:len(module)] == module {
            markup.InlineKeyboard = append(
                markup.InlineKeyboard,
                []models.InlineKeyboardButton{{
                    Text: group.Name,
                    CallbackData: "schedule:"+strconv.Itoa(group.Id)+":"+today,
                }},
            )
        }
    }

    _, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
        ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
        MessageID: update.CallbackQuery.Message.Message.ID,
        Text:      "Выберите группу",
        ReplyMarkup: &markup,
    })
    if err != nil {
        fmt.Println("ERROR", err)
    }
}
