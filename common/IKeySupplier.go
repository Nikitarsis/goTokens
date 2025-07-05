package common

// IKeySupplier - интерфейс для работы с ключами
type IKeySupplier interface {
	NewKey() Key
}