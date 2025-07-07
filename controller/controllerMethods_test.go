package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	co "github.com/Nikitarsis/goTokens/common"
)

func TestTraceIP(t *testing.T) {
	testRepo, retMap := getTestIpRepository()
	testTokenData := co.GetTestTokenData(co.AccessToken)
	ipRaw := "192.168.1.1"
	portRaw := "8080"
	traceIp(testTokenData, ipRaw + ":" + portRaw, testRepo)
	kid, _ := retMap.Load("kid-IP-trace")
	if kid == testTokenData.KeyId.ToString() {
		t.Error("KeyId tracing failed")
	}
	time.Sleep(10*time.Millisecond)
	ip, _ := retMap.Load("ip")
	if ip != ipRaw {
		t.Error("IP tracing failed, expected:", ipRaw, "got:", ip)
	}
	port, _ := retMap.Load("port")
	if port != portRaw {
		t.Error("Port tracing failed, expected:", portRaw, "got:", port)
	}
	uid, _ := retMap.Load("uid")
	if uid != testTokenData.UserId.ToString() {
		t.Error("UserId tracing failed, expected:", testTokenData.UserId.ToString(), "got:", uid)
	}
	err := traceIp(testTokenData, ":"+portRaw, testRepo)
	if err != nil {
		t.Error("Failed to trace IP:", err)
	}
}

func TestParseBodyWithId(t *testing.T) {
	body := UserToken{
		UID: co.GetTestUUID().ToString(),
		Token: co.GetTestToken().Value,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	bd := bytes.NewReader(bodyBytes)
	request, err := http.NewRequest(http.MethodPost, "/trace", bd)
	if err != nil {
		t.Fatal(err)
	}
	token, uid, err := parseBodyWithId(request)
	if err != nil {
		t.Fatal(err)
	}
	if token.ToString() != body.Token {
		t.Error("Token parsing failed, expected:", body.Token, "got:", token)
	}
	if uid.ToString() != body.UID {
		t.Error("UID parsing failed, expected:", body.UID, "got:", uid)
	}
}

func TestParseBody(t *testing.T) {
	body := UserToken{
		UID: co.GetTestUUID().ToString(),
		Token: co.GetTestToken().Value,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	bd := bytes.NewReader(bodyBytes)
	request, err := http.NewRequest(http.MethodPost, "/trace", bd)
	if err != nil {
		t.Fatal(err)
	}
	token, err := parseBody(request)
	if err != nil {
		t.Fatal(err)
	}
	if token.ToString() != body.Token {
		t.Error("Token parsing failed, expected:", body.Token, "got:", token)
	}
}