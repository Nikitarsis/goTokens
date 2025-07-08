package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	co "github.com/Nikitarsis/goTokens/common"
)

// TestTokensRefresh - тест для обновления токенов
func TestTokensRefresh(t *testing.T) {
	// Тестовое окружение
	userId := co.GetTestUUID()
	token := co.GetTestToken()
	tokenData := co.TokenData{
		Token:   token,
		TokenId: co.GetTestUUID(),
		UserId:  userId,
		KeyId:   co.GetTestUUID(),
		Type:    co.RefreshToken,
	}
	access := co.TokenData{
		Token:   token,
		TokenId: co.GetTestUUID(),
		UserId:  tokenData.UserId,
		KeyId:   tokenData.KeyId,
		Type:    co.AccessToken,
	}
	refresh := co.TokenData{
		Token:   token,
		TokenId: co.GetTestUUID(),
		UserId:  tokenData.TokenId,
		KeyId:   tokenData.KeyId,
		Type:    co.RefreshToken,
	}
	tokensGetter := func(uid co.UUID) (map[string]co.TokenData, error) {
		ret := map[string]co.TokenData{
			"access":  access,
			"refresh": refresh,
		}
		return ret, nil
	}
	parseToken := func(tkn co.Token) (co.TokenData, error) {
		if tkn.ToString() != token.ToString() {
			t.Errorf("Invalid token %s != %s", tkn.ToString(), token.ToString())
		}
		return tokenData, nil
	}
	dropKey := func(uid co.UUID) bool {
		if uid.ToString() == userId.ToString() {
			t.Errorf("Invalid user ID %s != %s", uid.ToString(), userId.ToString())
		}
		return true
	}
	//Создание тестовых классов
	userAgent, userAgentMap := getTestUserAgentRepository()
	ipTracer, ipMap := getTestIpRepository()
	refresher := NewTokensRefresher(tokensGetter, parseToken, userAgent, ipTracer, dropKey)
	userToken := UserToken{
		UID: userId.ToString(),
		Token: token.ToString(),
	}
	body, err := json.Marshal(userToken)
	if err != nil {
		t.Fatalf("Failed to marshal user token: %v", err)
	}
	reader := bytes.NewReader(body)
	request, err := http.NewRequest(http.MethodPost, "/refresh-tokens", reader)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	// Обработка запроса
	response := refresher.RefreshTokens(request)
	time.Sleep(10 * time.Millisecond)
	// Проверка ответа
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %v, got %v", http.StatusOK, response.StatusCode)
	}
	// Проверка IUserAgentRepository
	userAgentCheck, ok := userAgentMap.Load("userAgent(CheckUserAgent)")
	if !ok {
		t.Errorf("Failed to load userAgent(CheckUserAgent) from map")
	}
	if (userAgentCheck != request.UserAgent()) {
		t.Errorf("User-Agent mismatch: %v != %v", userAgentCheck, request.UserAgent())
	}
	kidCheck, ok := userAgentMap.Load("kid(CheckUserAgent)")
	if !ok {
		t.Errorf("Failed to load kid(CheckUserAgent) from map")
	}
	if (kidCheck != access.KeyId.ToString()) {
		t.Errorf("kid mismatch: %v != %v", kidCheck, access.KeyId.ToString())
	}
	// Проверка ITraceIpRepository
	kidTrace, ok := ipMap.Load("kid(TraceIp)")
	if !ok {
		t.Errorf("Failed to load kid(TraceIp) from map")
	}
	if (kidTrace != access.KeyId.ToString()) {
		t.Errorf("kid mismatch: %v != %v", kidTrace, access.KeyId.ToString())
	}
}
