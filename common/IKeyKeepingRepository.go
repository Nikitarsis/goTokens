package common

// IKeyKeepingRepository - интерфейс для работы с хранилищем ключей
type IKeyKeepingRepository interface {
	SaveKey(key Key)
	GetKey(kid UUID) (Key, bool)
}
