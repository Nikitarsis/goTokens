package tokenProducer

import (
	co "github.com/Nikitarsis/goTokens/common"
)

// TokenComponent - компонент для работы с данным модулем
type TokenComponent struct {
	producer          *tokenProducer
	checker           *tokenParser
	keyRepository     co.IKeyKeepingRepository
	componentSupplier *simpleComponentSupplier
}

// NewTokenComponentDefault создает новый экземпляр TokenComponent с конфигурацией
func NewTokenComponentDefault(keyRepository co.IKeyKeepingRepository, config ITokenComponentConfig) *TokenComponent {
	componentSupplier := newSimpleComponentSupplier(config.GetKeyChannelSize(), config.GetIdChannelSize())
	idSupply := componentSupplier.NewId
	idFactory := keyRepository.GetKey
	return &TokenComponent{
		producer:          NewTokenProducer(co.NewIssuer(config.GetIssuer()), idSupply),
		checker:           NewTokenParser(idFactory),
		keyRepository:     keyRepository,
		componentSupplier: componentSupplier,
	}
}

// CreateTokens создает пару токенов (access и refresh) для данного пользователя
func (tc *TokenComponent) CreateTokens(uid co.UUID) (map[string]co.TokenData, error) {
	key := tc.componentSupplier.NewKey()
	// Создание access токена
	access, err := tc.producer.CreateAccessToken(key, uid)
	if err != nil {
		return nil, err
	}
	// Создание refresh токена
	refresh, err := tc.producer.CreateRefreshToken(key, uid)
	if err != nil {
		return nil, err
	}
	// Асинхронное сохранение ключа
	go func() {
		tc.keyRepository.SaveKey(key)
	}()
	return map[string]co.TokenData{
		"access":  access,
		"refresh": refresh,
	}, nil
}
