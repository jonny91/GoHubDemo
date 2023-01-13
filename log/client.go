package log

import (
	"bytes"
	"fmt"
	"github.com/jonny91/hubDemo/registry"
	stlog "log"
	"net/http"
)

func SetClientLogger(serviceURL string, clientService registry.ServiceName) {
	stlog.SetPrefix(fmt.Sprintf("[%v] - ", clientService))
	stlog.SetFlags(0)
	stlog.SetOutput(&clientLogger{url: serviceURL})
}

type clientLogger struct {
	url string
}

func (cl *clientLogger) Write(p []byte) (n int, err error) {
	b := bytes.NewBuffer([]byte(p))
	res, err := http.Post(cl.url+"/log", "text/plain", b)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to send log message. service responded with code: %d", res.StatusCode)
	}
	return len(p), nil
}
