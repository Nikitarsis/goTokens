package common

type IKeyKeepingRepository interface {
	SaveKey(kid UUID, key Key) error
	GetKey(kid UUID) (Key, bool)
}