package bitrix_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/RB-PRO/SmartLogisterNotificationBot/pkg/bitrix"
)

// Звонки
func TestCrmLeadListCall(t *testing.T) {
	b, ErrNewB24 := bitrix.NewBitrix24("https://smartlogister.bitrix24.ru/rest/42/ewys918r6lz1q848/")
	if ErrNewB24 != nil {
		t.Error(ErrNewB24)
	}
	DateROI := time.Date(2023, time.June, 23, 0, 0, 0, 0, time.Local)

	ReqList := make(map[string]string)
	ReqList["filter[SOURCE_ID]"] = "CALL"
	// ReqList["filter[ASSIGNED_BY_ID]"] = "22"
	list, ErrCrmLeadList := b.CrmLeadList(DateROI, ReqList)
	if ErrCrmLeadList != nil {
		t.Error(ErrCrmLeadList)
	}
	fmt.Println("len(list.Result)", len(list.Result))
	// fmt.Println(list.Result[0].ID)
}

// Почта
func TestCrmLeadListInfo(t *testing.T) {
	b, ErrNewB24 := bitrix.NewBitrix24("https://smartlogister.bitrix24.ru/rest/42/ewys918r6lz1q848/")
	if ErrNewB24 != nil {
		t.Error(ErrNewB24)
	}
	DateROI := time.Date(2023, time.June, 23, 0, 0, 0, 0, time.Local)

	ReqList := make(map[string]string)
	ReqList["filter[SOURCE_ID]"] = "ADVERTISING"
	list, ErrCrmLeadList := b.CrmLeadList(DateROI, ReqList)
	if ErrCrmLeadList != nil {
		t.Error(ErrCrmLeadList)
	}
	fmt.Println("len(list.Result)", len(list.Result))
	// fmt.Println(list.Result[0].ID)
}

// Форма
func TestCrmLeadListForm(t *testing.T) {
	b, ErrNewB24 := bitrix.NewBitrix24("https://smartlogister.bitrix24.ru/rest/42/ewys918r6lz1q848/")
	if ErrNewB24 != nil {
		t.Error(ErrNewB24)
	}
	DateROI := time.Date(2023, time.June, 23, 0, 0, 0, 0, time.Local)

	ReqList := make(map[string]string)
	ReqList["filter[SOURCE_ID]"] = "WEB"
	list, ErrCrmLeadList := b.CrmLeadList(DateROI, ReqList)
	if ErrCrmLeadList != nil {
		t.Error(ErrCrmLeadList)
	}
	fmt.Println("len(list.Result)", len(list.Result))
	// fmt.Println(list.Result[0].ID)
}
