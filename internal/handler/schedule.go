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

func lesson_time(lesson eservice.Lesson) string {
    start, end, err := eservice.TimeByClass(lesson.Time)
    if err != nil {
        return ""
    }
    return start.Format("15:04")+"-"+end.Format("15:04")
}

func format_lesson(lesson eservice.Lesson) string {
    return strconv.Itoa(lesson.Time) + "(" + lesson_time(lesson) + ") " + lesson.TypeWork + " " + lesson.AuditCorps + 
           "\n"+lesson.Lesson[:len(lesson.Lesson)-len(lesson.TypeWork)]
}

func join_lessons(lessons []eservice.Lesson) string {
    if len(lessons) == 0 {
        return "Нет пар"
    }
    res := make([]string, len(lessons))
    for i, lesson := range lessons {
        res[i] = format_lesson(lesson)
    }
    return strings.Join(res, "\n\n")
}

func HandleSchedule(ctx context.Context, b *bot.Bot, update *models.Update) {
    group := strings.Split(update.CallbackQuery.Data, ":")[1]
    formated_date := strings.Split(update.CallbackQuery.Data, ":")[2]
    date, err := time.Parse("02.01.2006", formated_date)
    if err != nil {
        fmt.Println("ERROR", err)
    }

    markup := models.InlineKeyboardMarkup{}

    markup.InlineKeyboard = append(markup.InlineKeyboard, []models.InlineKeyboardButton{{
        Text: "<-",
        CallbackData: "schedule:" + group + ":" +date.Add(-24*time.Hour).Format("02.01.2006"),
    },
    {
        Text: "->",
        CallbackData: "schedule:" + group + ":" +date.Add(24*time.Hour).Format("02.01.2006"),
    }})

    schedule, err := eservice.Schedule(group)
    if err != nil {
        b.SendMessage(ctx,
            &bot.SendMessageParams{
                ChatID: update.Message.Chat.ID,
                Text: "Что то пошло не так...попробуй ещё раз",
            },
        )
        return
    }
    normalized_schedule := make(map[string][]eservice.Lesson)
    for _, record := range schedule {
        normalized_schedule[record.Day] = record.Lessons
    }
    lessons, ok := normalized_schedule[date.Format("02.01.2006")]
    if !ok {
        lessons = make([]eservice.Lesson, 0)
    }

    _, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
        ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
        MessageID: update.CallbackQuery.Message.Message.ID,
        Text:      formated_date + "\n\n" + join_lessons(lessons),
        ReplyMarkup: &markup,
    })
    if err != nil {
        fmt.Println("ERROR", err)
    }
}
