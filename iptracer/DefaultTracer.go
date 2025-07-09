package iptracer

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	co "github.com/Nikitarsis/goTokens/common"
)

// DefaultTracer реализует интерфейс IIpTracer
type DefaultTracer struct {
	url     string
	buffer  []string
	wg      sync.WaitGroup
	delay   time.Duration
	saveIp  func(co.DataIP) error
	checkIp func(co.DataIP) bool
}

// CreateDefaultTracer создает новый экземпляр DefaultTracer
func CreateDefaultTracer(config ITracerConfig, saveIp func(co.DataIP) error, checkIp func(co.DataIP) bool) co.IIpTracer {
	ret := &DefaultTracer{
		url:     config.GetWebhookURL(),
		buffer:  make([]string, config.GetBufferSize()),
		wg:      sync.WaitGroup{},
		delay:   config.GetDelay(),
		saveIp:  saveIp,
		checkIp: checkIp,
	}
	if config.ShouldSendWebhookMessage() {
		go ret.msgLoop()
	} else {
		go ret.logToStdLoop()
	}
	return ret
}

// sendMessage отправляет сообщение в вебхук
func (dt *DefaultTracer) sendMessage() {
	if len(dt.buffer) == 0 {
		return
	}
	dt.wg.Add(1)
	defer dt.wg.Done()
	_, err := http.Post(dt.url, "text/plain", strings.NewReader(strings.Join(dt.buffer, "\n")))
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	dt.buffer = dt.buffer[:0]
}

// msgLoop отправляет сообщения в вебхук с заданной периодичностью
func (dt *DefaultTracer) msgLoop() {
	for {
		dt.sendMessage()
		time.Sleep(dt.delay)
	}
}

// logToStdLoop отправляет сообщения в стандартный вывод с заданной периодичностью
func (dt *DefaultTracer) logToStdLoop() {
	for {
		fmt.Print(strings.Join(dt.buffer, "\n"))
		time.Sleep(dt.delay)
	}
}

// parseIpData формирует строку JSON из данных IP
func (dt *DefaultTracer) parseIpData(ip co.DataIP) string {
	ret := fmt.Sprintf("{\"kid\"=\"%s\", ", ip.KeyId.ToString())
	ret += fmt.Sprintf("\"uid\"=\"%s\", ", ip.UserId.ToString())
	ret += fmt.Sprintf("\"ip\"=\"%s:%d\"}", ip.IP.String(), ip.Port)
	return ret
}

// TraceIp отслеживает IP-адрес
func (dt *DefaultTracer) TraceIp(ip co.DataIP) {
	check := dt.checkIp(ip)
	if check {
		return
	}
	dt.saveIp(ip)
	dt.wg.Wait()
	dt.buffer = append(dt.buffer, dt.parseIpData(ip))
}
