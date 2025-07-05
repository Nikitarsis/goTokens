package common

type IKeyKeepingRepository interface {
	SaveKey(key Key)
	GetKey(kid UUID) (Key, bool)
}
