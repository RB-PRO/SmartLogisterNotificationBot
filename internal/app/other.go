package app

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

// Данные, которые необходимы для запуска приложения
type Authorization struct {
	ChatID            int64    `json:"ChatID"`            // ID чата, в который надо отправлять сообщение
	TelegramToken     string   `json:"TelegramToken"`     // Токен для телеграм бота
	BitrixToken       string   `json:"BitrixToken"`       // Ссылка на битрикс
	YandexDirectToken string   `json:"YandexDirectToken"` // Токен яндекс директа
	Companys          []string `json:"Companys"`          // Рекламные компании
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
