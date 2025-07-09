package interfaces

// IPostgresConfig определяет интерфейс конфигурации PostgreSQL
type IPostgresConfig interface {
	GetConnectionString() string
}
