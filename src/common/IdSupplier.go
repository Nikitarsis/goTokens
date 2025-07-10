package common

// IdSupplier - интерфейс для генерации идентификаторов
type IdSupplier interface {
	// NewId - генерирует новый уникальный идентификатор
	NewId() UUID
}
