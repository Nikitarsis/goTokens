package common

// IKeyKeepingRepository - интерфейс для работы с хранилищем ключей без функций удаления и изменения
type IKeyKeepingRepository interface {
	SaveKey(key Key)
	GetKey(kid UUID) (Key, bool)
}
