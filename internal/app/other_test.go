package app

import (
	"fmt"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	// Загрузить файл
	auf, ErrAuf := LoadAuthorization("authorization.json")
	if ErrAuf != nil {
		t.Error(ErrAuf)
	}

	// Создать приложение
	app, ErrApp := NewApplication(auf)
	if ErrApp != nil {
		t.Error(ErrApp)
	}

	Time1 := time.Now()
	Time1 = Time1.AddDate(0, 0, -1)
	// Time1 := time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC)
	Report1, ErrReport1 := app.Report1(Time1)
	if ErrReport1 != nil {
		t.Error(ErrReport1)
	}
	fmt.Println(Report1)
}
