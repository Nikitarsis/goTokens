package tokenProducer

import (
	"testing"

	co "github.com/Nikitarsis/goTokens/common"
)

// Тестовый парсер токенов
func getTestTokenParser(keyFunc func(co.UUID) (co.Key, bool)) *tokenParser {
	return NewTokenParser(keyFunc)
}

// TestParseToken создаёт тестирует парсинг токена
func TestParseToken(t *testing.T) {
	componentSupplier := getTestComponentSupplier()
	uid := componentSupplier.NewId()
	key := componentSupplier.NewKey()
	keyFunc := func(id co.UUID) (co.Key, bool) {
		if id == key.GetKid() {
			return key, true
		}
		t.Error("Failed to get key")
		return co.Key{}, false
	}
	producer := getTestTokenProducer(componentSupplier)
	parser := getTestTokenParser(keyFunc)
	_, token, _ := producer.createToken(key, uid, co.AccessToken)
	tokenData, err := parser.GetTokenData(token)
	if err != nil {
		t.Fatal(err)
	}
	if tokenData.Type != co.AccessToken {
		t.Error("Expected AccessToken, but got", tokenData.Type)
	}
	if tokenData.KeyId.ToString() != key.GetKid().ToString() {
		t.Error("Expected matching KeyId, but got", tokenData.KeyId.ToString())
	}
	if tokenData.UserId.ToString() != uid.ToString() {
		t.Error("Expected matching UserId, but got", tokenData.UserId.ToString())
	}
}
