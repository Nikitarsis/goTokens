package controller

// UserToken - структура для хранения токена пользователя и его UID
type UserToken struct {
	UID   string `json:"uid"`
	Token string `json:"token"`
}
