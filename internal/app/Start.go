package app

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron"
)

func Start() {
	// Загрузить файл
	auf, ErrAuf := LoadAuthorization("authorization.json")
	if ErrAuf != nil {
		panic(ErrAuf)
	}

	// Создать приложение
	app, ErrApp := NewApplication(auf)
	if ErrApp != nil {
		panic(ErrApp)
	}

	c := cron.New()

	// Отправка первой формы
	c.AddFunc("0 0 10 * *", func() {
		Time1 := time.Now()
		Time1 = Time1.AddDate(0, 0, -1)
		// Time1 := time.Date(2023, time.June, 23, 0, 0, 0, 0, time.UTC)
		Report1, ErrReport1 := app.Report1(Time1)
		if ErrReport1 != nil {
			log.Println(ErrReport1)
		} else {
			ErrSend := app.send(Time1, Report1)
			if ErrSend != nil {
				log.Println(ErrSend)
			}
		}
	})

	// Отправка второй формы
	c.AddFunc("0 0 11,15,18 * *", func() {
		Times := time.Now()
		Report2, ErrReport2 := app.Report2(Times)
		if ErrReport2 != nil {
			log.Println(ErrReport2)
		} else {
			ErrSend := app.send(Times, Report2)
			if ErrSend != nil {
				log.Println(ErrSend)
			}
		}
	})

	c.Start()
	select {}
}

func SendMail2() {
	fmt.Println("GOGOGO2")
}
