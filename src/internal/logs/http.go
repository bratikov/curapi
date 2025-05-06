package logs

import (
	"log"
)

type httpLogger struct{}

func New(prefix string) *log.Logger {
	return log.New(&httpLogger{}, prefix, 0)
}

func (l *httpLogger) Write(data []byte) (int, error) {
	Logger.Info().CallerSkipFrame(3).Msg(string(data))

	return len(data), nil
}
