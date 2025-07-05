package common

// IdSupplier - интерфейс для генерации идентификаторов
type IdSupplier interface {
	NewId() UUID
}
