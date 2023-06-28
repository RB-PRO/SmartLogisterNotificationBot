package app

import (
	"fmt"
	"os"

	"github.com/RB-PRO/SmartLogisterNotificationBot/pkg/bitrix"
	"github.com/nikoksr/notify/service/telegram"
	"github.com/zoer/yandex-api/direct"
)

// Структура приложения
type Application struct {
	B24           *bitrix.Bitrix24   // Работа с битриксом
	Authorization                    // Структура данных авторизации
	TG            *telegram.Telegram // Телеграм-нотификатор
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

	//
	token := os.Getenv(" ")
	client := direct.NewClient(token)
	//   client := direct.NewClient(" ")
	// Campaigns list
	campaigns, _ := client.Campaigns.GetList()
	fmt.Println(campaigns)
	// fmt.Println(len(campaigns), campaigns[0].Name, campaigns[0].Sum)

	return &Application{B24: b, TG: TG}, nil
}
