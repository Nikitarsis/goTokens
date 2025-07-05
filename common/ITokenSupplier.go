package common

type ITokenSupplier interface {
	NextToken() (UUID, Key)
}