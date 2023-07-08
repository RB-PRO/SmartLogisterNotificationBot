package app

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/RB-PRO/SmartLogisterNotificationBot/pkg/bitrix"
)

// Данные, которые необходимы для запуска приложения
type Authorization struct {
	ChatID            int64    `json:"ChatID"`            // ID чата, в который надо отправлять сообщение
	TelegramToken     string   `json:"TelegramToken"`     // Токен для телеграм бота
	BitrixToken       string   `json:"BitrixToken"`       // Ссылка на битрикс
	YandexDirectToken string   `json:"YandexDirectToken"` // Токен яндекс директа
	Companys          []string `json:"Companys"`          // Рекламные компании
	Companyss         []struct {
		Telegram string `json:"telegram"`
		Yandex   string `json:"yandex"`
		Bitrix   int    `json:"bitrix"`
	} `json:"Companyss"`
}

// Загрузить данные из файла
func LoadAuthorization(filename string) (prog Authorization, ErrorFIle error) {
	// Открыть файл
	jsonFile, ErrorFIle := os.Open(filename)
	if ErrorFIle != nil {
		return prog, ErrorFIle
	}
	defer jsonFile.Close()

	// Прочитать файл и получить массив byte
	jsonData, ErrorFIle := io.ReadAll(jsonFile)
	if ErrorFIle != nil {
		return prog, ErrorFIle
	}

	// Распарсить массив byte в структуру
	if ErrorFIle := json.Unmarshal(jsonData, &prog); ErrorFIle != nil {
		return prog, ErrorFIle
	}
	return prog, ErrorFIle
}

// Проверить данные на корректность
func (auf *Authorization) IsСorrect() (bool, error) {
	if auf.TelegramToken == "" {
		return false, errors.New("IsСorrect: Нет данных в auf.TelegramToken")
	}
	if auf.BitrixToken == "" {
		return false, errors.New("IsСorrect: Нет данных в auf.BitrixToken")
	}
	return true, nil
}

// Подстчитать к-во "некачественных" лидов в выгрузке
func CoutBad(list bitrix.CrmLeadListRes) (Total int) {

	// Цикл по всему слайсу лидов
	for _, lead := range list.Result {
		if IsBad(lead.StatusID) {
			Total++
		}
	}

	return Total
}

// Проверить, является ли лид некачественным по статусу
func IsBad(SourceID string) bool {
	// lead.SourceID == "1" || // 1) Дорого - качественный от 07.07.23
	switch SourceID {
	case "11": // 2) Хочет везти на наш контракт
		return true
	case "14": // 3) Карго
		return true
	case "12": // 4) Иной контрагент
		return true
	case "8": // 5) Нет доков на опасный груз
		return true
	case "5": // 6) Физ лицо
		return true
	case "9": // 7) На экспресс почту
		return true
	case "JUNK": // 8) Некачественный лид
		return true
	case "13": // 9) Спам
		return true
	case "7": // 10) Вес до 1 кг
		return true
	default:
		return false
	}
}

// Подстчитать к-во "качественных" лидов в выгрузке
func CoutGood(list bitrix.CrmLeadListRes) (Total int) {

	// Цикл по всему слайсу лидов
	for _, lead := range list.Result {
		if IsGood(lead.StatusID) {
			Total++
		}
	}

	return Total
}

// Проверка того, что лид корректен для учёта
func IsGood(SourceID string) bool {
	switch SourceID {
	case "IN_PROCESS": // Расчёт ставки
		return true
	case "PROCESSED": // КП отправлено
		return true
	case "1": // Созданы контакты
		return true
	default:
		return false
	}
}

// Подстчитать к-во "неопределённых" лидов в выгрузке
func CoutIndefinite(list bitrix.CrmLeadListRes) (Total int) {

	// Цикл по всему слайсу лидов
	for _, lead := range list.Result {
		if IsIndefinite(lead.StatusID) {
			Total++
		}
	}

	return Total
}

// Проверка того, что лид в неопределённом состоянии
func IsIndefinite(SourceID string) bool {
	switch SourceID {
	case "NEW": // Новый запрос
		return true
	case "2": // Менеджер взял ;))
		return true
	default:
		return false
	}
}

// Оставить с структуре только лиды определённого поля "Обращение за услугами"
//   - 2420 - Логистика
//   - 2422 - Таможенное оформление
//   - 2424 - Выкуп товара
func Filtered(leads bitrix.CrmLeadListRes, UF_CRM int) bitrix.CrmLeadListRes {
	var NewLeads bitrix.CrmLeadListRes
	for _, val := range leads.Result {
		if contains(val.UfCrm1688477669, UF_CRM) {
			NewLeads.Result = append(NewLeads.Result, val)
		}
	}
	NewLeads.Total = len(NewLeads.Result)
	return NewLeads
}

func contains(elems []int, v int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
