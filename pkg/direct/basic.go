// Пакет для взаимодействия с яндекс директ api

package direct

type YanDir struct {
	Token string // Токен от Yandex
}

// Создать новый клиент для работы с яндекс директом
func NewYandexDirectClient(Token string) *YanDir {
	return &YanDir{Token: Token}
}
