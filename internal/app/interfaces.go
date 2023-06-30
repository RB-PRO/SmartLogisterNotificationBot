package app

import (
	"fmt"

	"github.com/RB-PRO/SmartLogisterNotificationBot/pkg/bitrix"
	"github.com/RB-PRO/SmartLogisterNotificationBot/pkg/direct"
	"github.com/nikoksr/notify/service/telegram"
)

// Структура приложения
type Application struct {
	B24           *bitrix.Bitrix24   // Работа с битриксом
	Authorization                    // Структура данных авторизации
	TG            *telegram.Telegram // Телеграм-нотификатор
	YD            *direct.YanDir     // Яндекс-директ
	Companys      []string           // Все компании
}

// Создать приложение телеграм бота со всеми авторизациями
func NewApplication(auf Authorization) (*Application, error) {

	// Битрикс
	b, ErrNewB24 := bitrix.NewBitrix24(auf.BitrixToken)
	if ErrNewB24 != nil {
		return nil, fmt.Errorf("NewApplication: %v", ErrNewB24)
	}

	// Телеграм
	TG, ErrTG := NewNotification(auf.TelegramToken, auf.ChatID)
	if ErrTG != nil {
		return nil, fmt.Errorf("NewApplication: %v", ErrTG)
	}

	// Яндекс директ
	YD := direct.NewYandexDirectClient(auf.YandexDirectToken)

	return &Application{B24: b, TG: TG, YD: YD, Companys: auf.Companys}, nil
}
