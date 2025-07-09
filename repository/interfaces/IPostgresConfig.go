package interfaces

type IPostgresConfig interface {
	GetConnectionString() string
}
