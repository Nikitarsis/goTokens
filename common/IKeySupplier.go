package common

type IKeySupplier interface {
	NewKey() Key
}