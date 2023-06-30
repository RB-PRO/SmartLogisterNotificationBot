package direct

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func (yd *YanDir) ReportDo(DateFrom, DateTo time.Time) (string, error) {

	RawParam := `{
    "params": { 
    "SelectionCriteria": {
        "DateFrom": "%s",
        "DateTo": "%s"
    },
    "FieldNames": [ "Cost", "CampaignName"], 
 
    "ReportName": "Actual Data 123",
    "ReportType": "CAMPAIGN_PERFORMANCE_REPORT",
    "DateRangeType": "CUSTOM_DATE",
    "Format": "TSV",
    "IncludeVAT": "YES",
    "IncludeDiscount": "YES" 
    } 
}`

	payload := strings.NewReader(fmt.Sprintf(RawParam, DateFrom.Format("2006-01-02"), DateTo.Format("2006-01-02")))

	// Создаём запрос
	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, "https://api.direct.yandex.ru/json/v5/reports", payload)
	if ErrNewRequest != nil {
		return "", ErrNewRequest
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("processingMode", "online")
	req.Header.Add("skipReportHeader", "true")
	// req.Header.Add("returnMoneyInMicros", "false")
	req.Header.Add("skipReportSummary", "true")
	req.Header.Add("Authorization", "Bearer "+yd.Token)

	// Выполнить запрос
	res, ErrDo := client.Do(req)
	if ErrDo != nil {
		return "", ErrDo
	}
	defer res.Body.Close()

	// Читаем файл в слайс байтов
	body, ErrReadAll := io.ReadAll(res.Body)
	if ErrReadAll != nil {
		return "", ErrReadAll
	}

	// Создать файл
	FileName := "report.tsv"
	file, ErrCreateFile := os.Create(FileName)
	if ErrCreateFile != nil {
		return "", ErrCreateFile
	}

	// Внести данные в файл
	_, ErrWiriteFile := file.Write(body)
	if ErrWiriteFile != nil {
		return "", ErrWiriteFile
	}

	// Закрыть файл
	ErrClose := file.Close()
	if ErrClose != nil {
		return "", ErrClose
	}

	return FileName, nil
}
