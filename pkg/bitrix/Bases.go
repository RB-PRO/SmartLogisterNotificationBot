package bitrix

import (
	"errors"
	"net/http"
)

// Структура битрикс приложения
type Bitrix24 struct {
	URL    string // Ссылка на авторизованное приложение на входящий вебхук
	Client *http.Client
}

func NewBitrix24(URL string) (*Bitrix24, error) {
	if URL == "" {
		return nil, errors.New("NewBitrix24: nil of URL")
	}
	client := &http.Client{}
	return &Bitrix24{URL: URL, Client: client}, nil
}
