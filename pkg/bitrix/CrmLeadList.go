package bitrix

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Название метода
const CrmLeadList string = "crm.lead.list"

// Структура ответа к серверу
type CrmLeadListReq struct {
	DATE_CREATE_OT string `json:">DATE_CREATE"`
	DATE_CREATE_DO string `json:"<DATE_CREATE"`
}

// time.Now().Format("2006-01-02T15:04:05")
func (b *Bitrix24) CrmLeadList(DateROI time.Time, RequestsList map[string]string) (CrmLeadListRes, error) {

	// Формирование определённого запроса
	var ReqStruct CrmLeadListReq
	ReqStruct.DATE_CREATE_OT = time.Date(DateROI.Year(), DateROI.Month(), DateROI.Day(), 0, 0, 0, 0, time.UTC).Format("2006-01-02T15:04:05")
	ReqStruct.DATE_CREATE_DO = time.Date(DateROI.Year(), DateROI.Month(), DateROI.Day(), 23, 59, 0, 0, time.UTC).Format("2006-01-02T15:04:05")
	fmt.Printf("TIME: %+#v\n", ReqStruct)

	// Выполнить запрос
	Request, ErrRequest := http.NewRequest(http.MethodGet, b.URL+CrmLeadList, nil)
	if ErrRequest != nil {
		return CrmLeadListRes{}, ErrRequest
	}

	q := Request.URL.Query()
	q.Add("filter[>DATE_CREATE]", ReqStruct.DATE_CREATE_OT) // Дата создания от
	q.Add("filter[<DATE_CREATE]", ReqStruct.DATE_CREATE_DO) // Дата создания до
	q.Add("filter[!STATUS_ID][0]", "4")                     // Игнирируем статус "Прочее"
	q.Add("filter[!STATUS_ID][1]", "12")                    // Игнирируем статус "Иной контрагент"

	// Добавляем все остальные структуры
	for key, val := range RequestsList {
		q.Add(key, val)
	}

	Request.URL.RawQuery = q.Encode()
	fmt.Println("Do URL:", Request.URL.String())

	// Выполнить запрос
	Response, ErrDo := b.Client.Do(Request)
	if ErrDo != nil {
		return CrmLeadListRes{}, ErrDo
	}
	defer Response.Body.Close()

	// Получить массив байтов
	ResponseBody, ErrReadAll := io.ReadAll(Response.Body)
	if ErrReadAll != nil {
		return CrmLeadListRes{}, ErrReadAll
	}

	// Распарсить массив byte в структуру
	var ResStruct CrmLeadListRes
	ErrUnmarshal := json.Unmarshal(ResponseBody, &ResStruct)
	if ErrUnmarshal != nil {
		return CrmLeadListRes{}, ErrUnmarshal
	}

	return ResStruct, nil
}

// Структура ответа от сервера
type CrmLeadListRes struct {
	Result []struct {
		ID                  string    `json:"ID,omitempty"`
		Title               string    `json:"TITLE,omitempty"`
		Honorific           any       `json:"HONORIFIC,omitempty"`
		Name                string    `json:"NAME,omitempty"`
		SecondName          any       `json:"SECOND_NAME,omitempty"`
		LastName            any       `json:"LAST_NAME,omitempty"`
		CompanyTitle        any       `json:"COMPANY_TITLE,omitempty"`
		CompanyID           any       `json:"COMPANY_ID,omitempty"`
		ContactID           any       `json:"CONTACT_ID,omitempty"`
		IsReturnCustomer    string    `json:"IS_RETURN_CUSTOMER,omitempty"`
		Birthdate           string    `json:"BIRTHDATE,omitempty"`
		SourceID            string    `json:"SOURCE_ID,omitempty"`
		SourceDescription   string    `json:"SOURCE_DESCRIPTION,omitempty"`
		StatusID            string    `json:"STATUS_ID,omitempty"`
		StatusDescription   any       `json:"STATUS_DESCRIPTION,omitempty"`
		Post                any       `json:"POST,omitempty"`
		Comments            any       `json:"COMMENTS,omitempty"`
		CurrencyID          string    `json:"CURRENCY_ID,omitempty"`
		Opportunity         string    `json:"OPPORTUNITY,omitempty"`
		IsManualOpportunity string    `json:"IS_MANUAL_OPPORTUNITY,omitempty"`
		HasPhone            string    `json:"HAS_PHONE,omitempty"`
		HasEmail            string    `json:"HAS_EMAIL,omitempty"`
		HasImol             string    `json:"HAS_IMOL,omitempty"`
		AssignedByID        string    `json:"ASSIGNED_BY_ID,omitempty"`
		CreatedByID         string    `json:"CREATED_BY_ID,omitempty"`
		ModifyByID          string    `json:"MODIFY_BY_ID,omitempty"`
		DateCreate          time.Time `json:"DATE_CREATE,omitempty"`
		DateModify          time.Time `json:"DATE_MODIFY,omitempty"`
		DateClosed          string    `json:"DATE_CLOSED,omitempty"`
		StatusSemanticID    string    `json:"STATUS_SEMANTIC_ID,omitempty"`
		Opened              string    `json:"OPENED,omitempty"`
		OriginatorID        any       `json:"ORIGINATOR_ID,omitempty"`
		OriginID            any       `json:"ORIGIN_ID,omitempty"`
		MovedByID           string    `json:"MOVED_BY_ID,omitempty"`
		MovedTime           time.Time `json:"MOVED_TIME,omitempty"`
		Address             any       `json:"ADDRESS,omitempty"`
		Address2            any       `json:"ADDRESS_2,omitempty"`
		AddressCity         any       `json:"ADDRESS_CITY,omitempty"`
		AddressPostalCode   any       `json:"ADDRESS_POSTAL_CODE,omitempty"`
		AddressRegion       any       `json:"ADDRESS_REGION,omitempty"`
		AddressProvince     any       `json:"ADDRESS_PROVINCE,omitempty"`
		AddressCountry      any       `json:"ADDRESS_COUNTRY,omitempty"`
		AddressCountryCode  any       `json:"ADDRESS_COUNTRY_CODE,omitempty"`
		AddressLocAddrID    any       `json:"ADDRESS_LOC_ADDR_ID,omitempty"`
		UtmSource           any       `json:"UTM_SOURCE,omitempty"`
		UtmMedium           any       `json:"UTM_MEDIUM,omitempty"`
		UtmCampaign         any       `json:"UTM_CAMPAIGN,omitempty"`
		UtmContent          any       `json:"UTM_CONTENT,omitempty"`
		UtmTerm             any       `json:"UTM_TERM,omitempty"`
		ParentID137         any       `json:"PARENT_ID_137,omitempty"`
		LastActivityBy      string    `json:"LAST_ACTIVITY_BY,omitempty"`
		LastActivityTime    time.Time `json:"LAST_ACTIVITY_TIME,omitempty"`
	} `json:"result,omitempty"`
	Total int `json:"total,omitempty"`
	// Time  struct {
	// 	Start            float64   `json:"start,omitempty"`
	// 	Finish           float64   `json:"finish,omitempty"`
	// 	Duration         float64   `json:"duration,omitempty"`
	// 	Processing       float64   `json:"processing,omitempty"`
	// 	DateStart        time.Time `json:"date_start,omitempty"`
	// 	DateFinish       time.Time `json:"date_finish,omitempty"`
	// 	OperatingResetAt int       `json:"operating_reset_at,omitempty"`
	// 	Operating        int       `json:"operating,omitempty"`
	// } `json:"time,omitempty"`
}
