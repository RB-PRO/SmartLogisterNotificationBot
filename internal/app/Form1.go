package app

import (
	"fmt"
	"time"
)

// Отправить отчёт по первой форме
func (app *Application) Report1(DataDate time.Time) (string, error) {

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
