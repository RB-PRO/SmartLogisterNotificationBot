package app

import (
	"context"
	"time"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

func (app *Application) send(TecalTime time.Time, Message string) error {
	return app.TG.Send(
		context.Background(),
		"Отчёт за "+TecalTime.Format("15:04 02:01:2006"),
		Message,
	)
}

func NewNotification(token string, ChatID int64) (*telegram.Telegram, error) {

	// Create a telegram service. Ignoring error for demo simplicity.
	telegramService, ErrorServece := telegram.New(token)
	if ErrorServece != nil {
		return nil, ErrorServece
	}

	// Добавить ID,куда будут посылаться уведомления
	telegramService.AddReceivers(ChatID)

	// Tell our notifier to use the telegram service. You can repeat the above process
	// for as many services as you like and just tell the notifier to use them.
	// Inspired by http middlewares used in higher level libraries.
	notify.UseServices(telegramService)

	// Send a test message.
	ErrorTelegramSend := notify.Send(
		context.Background(),
		"Время "+time.Now().Format("15:04 02.01.2006"),
		"SmartLogisterNotification: Начинаю работу",
	)
	if ErrorTelegramSend != nil {
		return nil, ErrorTelegramSend
	}

	return telegramService, nil
}
