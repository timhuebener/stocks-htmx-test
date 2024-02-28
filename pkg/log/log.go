package log

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=log.go -destination=mocks/logger_mocks.go -package=loggermocks Logger

type Level uint8

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var lvlToString = map[Level]string{
	DEBUG:   "DEBUG",
	INFO:    "INFO",
	WARNING: "WARNING",
	ERROR:   "ERROR",
	FATAL:   "FATAL",
}

type Logger interface {
	Log(ctx context.Context, lvl Level, msg string, meta map[string]string)
	Debug(ctx context.Context, msg string, meta map[string]string)
	Info(ctx context.Context, msg string, meta map[string]string)
	Warning(ctx context.Context, msg string, meta map[string]string)
	Error(ctx context.Context, msg string, meta map[string]string)
	Fatal(ctx context.Context, msg string, meta map[string]string)
}

type Exporter interface {
	Export(line map[string]string)
}
