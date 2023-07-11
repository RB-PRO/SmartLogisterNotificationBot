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
	q.Add("SELECT[0]", "STATUS_ID")                         //
	q.Add("SELECT[1]", "SOURCE_ID")                         // Статус из справочника.
	q.Add("SELECT[2]", "UF_CRM_1688477669")                 // Пользовательское поле "Обращение за услугами"
	q.Add("SELECT[3]", "NAME")                              // Пользовательское поле "Обращение за услугами"

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
		ID                  string `json:"ID"`
		Title               string `json:"TITLE"`
		Honorific           any    `json:"HONORIFIC"`
		Name                any    `json:"NAME"`
		SecondName          any    `json:"SECOND_NAME"`
		LastName            any    `json:"LAST_NAME"`
		CompanyTitle        any    `json:"COMPANY_TITLE"`
		CompanyID           any    `json:"COMPANY_ID"`
		ContactID           any    `json:"CONTACT_ID"`
		IsReturnCustomer    string `json:"IS_RETURN_CUSTOMER"`
		Birthdate           string `json:"BIRTHDATE"`
		SourceID            string `json:"SOURCE_ID"`
		SourceDescription   string `json:"SOURCE_DESCRIPTION"`
		StatusID            string `json:"STATUS_ID"`
		StatusDescription   any    `json:"STATUS_DESCRIPTION"`
		Post                any    `json:"POST"`
		Comments            any    `json:"COMMENTS"`
		CurrencyID          string `json:"CURRENCY_ID"`
		Opportunity         string `json:"OPPORTUNITY"`
		IsManualOpportunity string `json:"IS_MANUAL_OPPORTUNITY"`
		HasPhone            string `json:"HAS_PHONE"`
		HasEmail            string `json:"HAS_EMAIL"`
		HasImol             string `json:"HAS_IMOL"`
		AssignedByID        string `json:"ASSIGNED_BY_ID"`
		CreatedByID         string `json:"CREATED_BY_ID"`
		ModifyByID          string `json:"MODIFY_BY_ID"`
		// DateCreate          time.Time `json:"DATE_CREATE"`
		// DateModify          time.Time `json:"DATE_MODIFY"`
		// DateClosed          time.Time `json:"DATE_CLOSED"`
		StatusSemanticID string `json:"STATUS_SEMANTIC_ID"`
		Opened           string `json:"OPENED"`
		OriginatorID     any    `json:"ORIGINATOR_ID"`
		OriginID         any    `json:"ORIGIN_ID"`
		MovedByID        string `json:"MOVED_BY_ID"`
		// MovedTime           time.Time `json:"MOVED_TIME"`
		Address            any    `json:"ADDRESS"`
		Address2           any    `json:"ADDRESS_2"`
		AddressCity        any    `json:"ADDRESS_CITY"`
		AddressPostalCode  any    `json:"ADDRESS_POSTAL_CODE"`
		AddressRegion      any    `json:"ADDRESS_REGION"`
		AddressProvince    any    `json:"ADDRESS_PROVINCE"`
		AddressCountry     any    `json:"ADDRESS_COUNTRY"`
		AddressCountryCode any    `json:"ADDRESS_COUNTRY_CODE"`
		AddressLocAddrID   any    `json:"ADDRESS_LOC_ADDR_ID"`
		UtmSource          string `json:"UTM_SOURCE"`
		UtmMedium          string `json:"UTM_MEDIUM"`
		UtmCampaign        string `json:"UTM_CAMPAIGN"`
		UtmContent         string `json:"UTM_CONTENT"`
		UtmTerm            string `json:"UTM_TERM"`
		ParentID137        any    `json:"PARENT_ID_137"`
		LastActivityBy     string `json:"LAST_ACTIVITY_BY"`
		// LastActivityTime    time.Time `json:"LAST_ACTIVITY_TIME"`
		UfCrm1688477669 interface{} `json:"UF_CRM_1688477669"`
	} `json:"result"`
	Total int `json:"total"`
	// Time  struct {
	// 	Start            float64   `json:"start"`
	// 	Finish           float64   `json:"finish"`
	// 	Duration         float64   `json:"duration"`
	// 	Processing       float64   `json:"processing"`
	// 	DateStart        time.Time `json:"date_start"`
	// 	DateFinish       time.Time `json:"date_finish"`
	// 	OperatingResetAt int       `json:"operating_reset_at"`
	// 	Operating        float64   `json:"operating"`
	// } `json:"time"`
}
