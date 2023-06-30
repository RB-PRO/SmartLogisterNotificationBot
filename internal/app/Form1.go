package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/RB-PRO/SmartLogisterNotificationBot/pkg/direct"
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

	// Яндекс запросы

	// Получение денных за месяц
	DateFrom := time.Date(DataDate.Year(), DataDate.Month(), 1, 0, 0, 0, 0, DataDate.Location())
	DateTo := time.Date(DataDate.Year(), DataDate.Month(), LastDayOfMonth(DataDate), 0, 0, 0, 0, DataDate.Location())
	FileNameMouth, ErrYandexReport := app.YD.ReportDo(DateFrom, DateTo)
	if ErrYandexReport != nil {
		return "", ErrYandexReport
	}
	DirectMouth, ErrUnwrapTSV := direct.UnwrapTSV(FileNameMouth)
	if ErrUnwrapTSV != nil {
		return "", ErrUnwrapTSV
	}

	// Получение денных за день
	FileNameDay, ErrYandexReport := app.YD.ReportDo(DataDate, DataDate)
	if ErrYandexReport != nil {
		return "", ErrYandexReport
	}
	DirectDay, ErrUnwrapTSV := direct.UnwrapTSV(FileNameDay)
	if ErrUnwrapTSV != nil {
		return "", ErrUnwrapTSV
	}

	///////////////////////////////////////////////////////////////////////////////////////////////////

	// За 23.06  | #лиды2023

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
	// Расход за мес = 48 091 руб

	// Формирование сообщения
	var Message string
	Message += fmt.Sprintf("За %v  | #лиды%d\n\n", DataDate.Format("02.01"), DataDate.Year())
	Message += "Логистика:\n"
	Message += fmt.Sprintf("- Формы: %d (%d)\n", LeadsForm.Total, IsGood(LeadsForm))
	Message += fmt.Sprintf("- Звонки: %d (%d)\n", LeadsCall.Total, IsGood(LeadsCall))
	Message += fmt.Sprintf("- Инфо: %d (%d)\n", LeadsInfo.Total, IsGood(LeadsInfo))
	Message += fmt.Sprintf("ИТОГО: %d (%d)\n", LeadsForm.Total+LeadsCall.Total+LeadsInfo.Total, IsGood(LeadsForm)+IsGood(LeadsCall)+IsGood(LeadsInfo))
	var Other float64
	for _, company := range app.Companys {
		Other += SumCoast(DirectDay, company)
	}
	AllRashodDay := SumCoast(DirectDay, "")
	Rashod := AllRashodDay - Other
	Message += fmt.Sprintf("Расход = %.0f руб (без учета ТО и Выкупа)\n", Rashod)
	Message += fmt.Sprintf("Стоимость 1 лида = %.0f руб\n", Rashod/float64(LeadsForm.Total+LeadsCall.Total+LeadsInfo.Total))
	Message += fmt.Sprintf("Общий тек. расход за %s = %.0f руб\n", DataDate.Format("02.01"), SumCoast(DirectMouth, ""))
	for _, company := range app.Companys {
		Message += "- - - - - - - - - - - - - - - -\n"
		Message += company + ":\n"
		Message += fmt.Sprintf("- Формы: %d (%d)\n", LeadsForm.Total, IsGood(LeadsForm))
		Message += fmt.Sprintf("- Звонки: %d (%d)\n", LeadsCall.Total, IsGood(LeadsCall))
		Message += fmt.Sprintf("- Инфо: %d (%d)\n", LeadsInfo.Total, IsGood(LeadsInfo))
		Message += fmt.Sprintf("ИТОГО: %d (%d)\n", LeadsForm.Total+LeadsCall.Total+LeadsInfo.Total, IsGood(LeadsForm)+IsGood(LeadsCall)+IsGood(LeadsInfo))

		Message += fmt.Sprintf("Расход всего = %.0f руб\n", SumCoast(DirectDay, company))
		Message += fmt.Sprintf("Расход за мес = %.0f руб\n", SumCoast(DirectMouth, company))
	}
	return Message, nil
}

func LastDayOfMonth(t time.Time) int {
	firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, 0).Add(-time.Nanosecond)
	return lastDay.Day()
}

// Подсчёт суммы по всем компаниям
func SumCoast(companys []direct.AnswerTSV, query string) (sum float64) {
	for _, company := range companys {
		if query == "" {
			sum += float64(company.Cost) / 1000000
		} else {
			if strings.Contains(strings.ToLower(company.CampaignName), query) {
				sum += float64(company.Cost) / 1000000
			}
		}
	}
	return sum
}
