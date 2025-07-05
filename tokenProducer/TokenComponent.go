package tokenProducer

import (
	co "github.com/Nikitarsis/goTokens/common"
)

type TokenComponent struct {
	producer          *tokenProducer
	checker           *tokenParser
	keyRepository     co.IKeyKeepingRepository
	componentSupplier *SimpleComponentSupplier
}

func NewTokenComponentDefault(keyRepository co.IKeyKeepingRepository, config ITokenComponentConfig) *TokenComponent {
	componentSupplier := NewSimpleComponentSupplier(config.GetKeyChannelSize(), config.GetJtiChannelSize())
	idSupply := componentSupplier.NewId
	idFactory := keyRepository.GetKey
	return &TokenComponent{
		producer:          NewTokenProducer(co.NewIssuer(config.GetIssuer()), idSupply),
		checker:           NewTokenParser(idFactory),
		keyRepository:     keyRepository,
		componentSupplier: componentSupplier,
	}
}

func (tc *TokenComponent) CreateTokens(uid co.UUID) (co.TokensPair, error) {
	key := tc.componentSupplier.NewKey()
	_, access, err := tc.producer.CreateAccessToken(key, uid)
	if err != nil {
		return co.TokensPair{}, err
	}
	_, refresh, err := tc.producer.CreateRefreshToken(key, uid)
	if err != nil {
		return co.TokensPair{}, err
	}
	return co.TokensPair{
		Access:  access,
		Refresh: refresh,
	}, nil
}
