package common

type ITokenSupplier interface {
	NextToken() Key
}