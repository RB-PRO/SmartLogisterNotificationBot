package app

import (
	"fmt"
	"time"

	"github.com/RB-PRO/SmartLogisterNotificationBot/pkg/bitrix"
)

// Отправить отчёт по второй форме
func (app *Application) Report2(DataDate time.Time) (string, error) {

	// Лиды звонков
	ReqCall := make(map[string]string)
	ReqCall["filter[SOURCE_ID]"] = "CALL"
	ReqCall["filter[ASSIGNED_BY_ID]"] = "22"
	LeadsCall, ErrCrmLeadListCall := app.B24.CrmLeadList(DataDate, ReqCall)
	if ErrCrmLeadListCall != nil {
		return "", ErrCrmLeadListCall
	}

	// Лиды сообщений почты
	ReqInfo := make(map[string]string)
	ReqInfo["filter[SOURCE_ID]"] = "ADVERTISING"
	LeadsInfo, ErrCrmLeadListInfo := app.B24.CrmLeadList(DataDate, ReqInfo)
	if ErrCrmLeadListInfo != nil {
		return "", ErrCrmLeadListInfo
	}

	// Лиды форм
	ReqForm := make(map[string]string)
	ReqForm["filter[SOURCE_ID]"] = "WEB"
	LeadsForm, ErrCrmLeadListForm := app.B24.CrmLeadList(DataDate, ReqForm)
	if ErrCrmLeadListForm != nil {
		return "", ErrCrmLeadListForm
	}

	// Формирование сообщения
	var Message string
	Message += fmt.Sprintf("За %v  | #лиды2023\n\n", DataDate.Format("02.01"))
	Message += "Логистика:\n"
	Message += fmt.Sprintf("- Формы: %d (%d)\n", LeadsForm.Total, IsGood(LeadsForm))
	Message += fmt.Sprintf("- Звонки: %d (%d)\n", LeadsCall.Total, IsGood(LeadsCall))
	Message += fmt.Sprintf("- Инфо: %d (%d)\n", LeadsInfo.Total, IsGood(LeadsInfo))
	Message += fmt.Sprintf("- ИТОГО: %d (%d)\n", LeadsForm.Total+LeadsCall.Total+LeadsInfo.Total, IsGood(LeadsForm)+IsGood(LeadsCall)+IsGood(LeadsInfo))

	// 	Message := `За %v  | #лиды2023

	// Логистика:
	// - Формы: 6 (5)
	// - Звонки: 3 (2)
	// - Инфо: 2 (2)
	// ИТОГО: 11 (9)
	// Расход = 23 673 руб (без учета ТО и Выкупа)
	// Стоимость 1 лида = 2 152 руб
	// Общий тек. расход за 06.23 = 642 373 руб
	// - - - - - - - - - - - - - - - -
	// Таможенное оформление:
	// - Формы: 0 (0)
	// - Звонки: 0 (0)
	// - Инфо: 0 (0)
	// ИТОГО: 0 (0)
	// Расход всего = 5 031 руб
	// Расход за мес = 58 185 руб
	// - - - - - - - - - - - - - - - -
	// Выкуп товара:
	// - Формы: 3 (2)
	// - Звонки: 0 (0)
	// - Инфо: 0 (0)
	// ИТОГО: 3 (2)
	// Расход = 4 060 руб
	// Расход за мес = 48 091 руб`

	return Message, nil
}

// Подстчитать к-во "хороших" лидов
func IsGood(list bitrix.CrmLeadListRes) (Total int) {

	// Цикл по всему слайсу лидов
	for _, lead := range list.Result {
		if lead.SourceID == "1" || // 1) Дорого
			lead.SourceID == "11" || // 2) Хочет везти на наш контракт
			lead.SourceID == "14" || // 3) Карго
			lead.SourceID == "12" || // 4) Иной контрагент
			lead.SourceID == "8" || // 5) Нет доков на опасный груз
			lead.SourceID == "5" || // 6) Физ лицо
			lead.SourceID == "9" || // 7) На экспресс почту
			lead.SourceID == "JUNK" || // 8) Некачественный лид
			lead.SourceID == "13" || // 9) Спам
			lead.SourceID == "7" { // 10) Вес до 1 кг
			continue
		}
		Total++
	}

	return Total
}
