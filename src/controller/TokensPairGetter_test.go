package controller

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	co "github.com/Nikitarsis/goTokens/common"
)

// Случайный токен
func getRandomTokenData() co.TokenData {
	return co.TokenData{
		Token:   co.Token{Value: "just.random.token"},
		TokenId: co.GetTestUUID(),
		UserId:  co.GetTestUUID(),
		KeyId:   co.GetTestUUID(),
		Type:    co.AccessToken,
	}
}

// Проверка результата выполнения внедрённых функций
func checkMap(
	translator func(int, string) string,
	kid co.UUID,
	userAgent string,
	ip string,
) (string, bool) {
	if translator(1, "userAgent(SaveUserAgent)") != "" {
		return "this method shouldn't be called", false
	}
	if translator(0, "ip(TraceIp)") != ip {
		return "Incorrect IP", false
	}
	if translator(0, "kid(TraceIp)") != kid.ToString() {
		ret := fmt.Sprintf("Incorrect IP kid %s, should be: %s", translator(0, "kid(TraceIp)"), kid.ToString())
		return ret, false
	}
	if translator(1, "kid(SaveUserAgent)") != kid.ToString() {
		return "Incorrect userAgent save kid", false
	}
	if translator(1, "userAgent(CheckUserAgent)") != userAgent {
		return "Incorrect userAgent", false
	}
	return "Ok", true
}

// TestRequestGet - тест для GET-запроса
func TestRequestGet(t *testing.T) {
	uid := co.GetTestUUID()
	reader := strings.NewReader("nothing")
	access := getRandomTokenData()
	refresh := getRandomTokenData()
	tokenPairProducer := func(id co.UUID) (map[string]co.TokenData, error) {
		if uid.ToString() != id.ToString() {
			t.Fatal("User ID mismatch")
		}
		return map[string]co.TokenData{
			"access":  access,
			"refresh": refresh,
		}, nil
	}
	url := "/tokens/?uid=" + strings.ReplaceAll(uid.ToString(), "+", "%2b")
	request, err := http.NewRequest(http.MethodGet, url, reader)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	controller, _ := getTestTokenPairGetter(tokenPairProducer)
	ret, err := controller.parseRequestGet(request)
	if err != nil {
		t.Fatalf("Failed to parse request: %v", err)
	}
	if ret != uid {
		t.Fatalf("Expected user ID %v, got %v", uid, ret)
	}
}

// TestRequestPost - тест для POST-запроса
func TestRequestPost(t *testing.T) {
	uid := co.GetTestUUID()
	body := fmt.Sprintf("{\"uid\":\"%s\"}", uid.ToString())
	reader := strings.NewReader(body)
	access := getRandomTokenData()
	refresh := getRandomTokenData()
	tokenPairProducer := func(id co.UUID) (map[string]co.TokenData, error) {
		if uid.ToString() != id.ToString() {
			t.Fatal("User ID mismatch")
		}
		return map[string]co.TokenData{
			"access":  access,
			"refresh": refresh,
		}, nil
	}
	url := "/tokens/"
	request, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	controller, _ := getTestTokenPairGetter(tokenPairProducer)
	ret, err := controller.parseRequestPost(request)
	if err != nil {
		t.Fatalf("Failed to parse request: %v", err)
	}
	if ret != uid {
		t.Fatalf("Expected user ID %v, got %v", uid, ret)
	}
}

// TestGetTokensPair - тест для получения пары токенов
func TestGetTokensPair(t *testing.T) {
	uid := co.GetTestUUID()
	reader := strings.NewReader("nothing")
	access := getRandomTokenData()
	refresh := getRandomTokenData()
	tokenPairProducer := func(id co.UUID) (map[string]co.TokenData, error) {
		if uid.ToString() != id.ToString() {
			t.Fatal("User ID mismatch")
		}
		return map[string]co.TokenData{
			"access":  access,
			"refresh": refresh,
		}, nil
	}
	url := "/tokens/?uid=" + strings.ReplaceAll(uid.ToString(), "+", "%2b")
	request, err := http.NewRequest(http.MethodGet, url, reader)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	controller, translated := getTestTokenPairGetter(tokenPairProducer)
	response := controller.GetTokensPair(request)
	if err != nil {
		t.Fatalf("Failed to parse request: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %v, got %v", http.StatusOK, response.StatusCode)
	}
	time.Sleep(100 * time.Millisecond)
	str, check := checkMap(translated, refresh.KeyId, request.UserAgent(), "")
	if !check {
		t.Fatalf("Map check failed: %v", str)
	}
}

