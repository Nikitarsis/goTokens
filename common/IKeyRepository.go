package common

type IKeyRepository interface {
	SaveKey(kid UUID, key Key) error
	GetKey(kid UUID) (Key, bool)
}