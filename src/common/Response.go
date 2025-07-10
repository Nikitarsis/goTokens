package common

// Response - структура для хранения ответа
type Response struct {
	// StatusCode - код статуса ответа
	StatusCode  int
	// Message - сообщение ответа
	Message     []byte
}
