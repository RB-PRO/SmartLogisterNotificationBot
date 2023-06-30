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
	Message := `Результаты на %v:

	%d форм
	%d звонка
	%d инфо

	ИТОГО: %d`
	Message = fmt.Sprintf(Message, DataDate.Format("15:04 02.01.2006"), LeadsForm.Total, LeadsCall.Total, LeadsInfo.Total, LeadsForm.Total+LeadsCall.Total+LeadsInfo.Total)

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
