package common

type IdSupplier interface {
	NewId() UUID
}
