package log

import (
	"context"
	"regexp"
	"time"
)

var _ Logger = &StdLogger{}

type StdLogger struct {
	rsc       string
	lvl       Level
	exp       Exporter
	hashVins  bool
	vinRegexp *regexp.Regexp
}

func NewStdLogger(resource string, lvl Level, exporter Exporter) StdLogger {
	return StdLogger{
		rsc: resource,
		lvl: lvl,
		exp: exporter,
	}
}

func (s StdLogger) Log(ctx context.Context, lvl Level, msg string, meta map[string]string) {
	if lvl < s.lvl {
		return
	}
	if meta == nil {
		meta = map[string]string{}
	}

	meta["timestamp"] = time.Now().UTC().Format(time.RFC3339)
	meta["service.name"] = s.rsc
	meta["lvl"] = lvlToString[lvl]
	meta["msg"] = msg

	s.exp.Export(meta)
}

func (s StdLogger) Debug(ctx context.Context, msg string, meta map[string]string) {
	s.Log(ctx, DEBUG, msg, meta)
}

func (s StdLogger) Info(ctx context.Context, msg string, meta map[string]string) {
	s.Log(ctx, INFO, msg, meta)
}

func (s StdLogger) Warning(ctx context.Context, msg string, meta map[string]string) {
	s.Log(ctx, WARNING, msg, meta)
}

func (s StdLogger) Error(ctx context.Context, msg string, meta map[string]string) {
	s.Log(ctx, ERROR, msg, meta)
}

func (s StdLogger) Fatal(ctx context.Context, msg string, meta map[string]string) {
	s.Log(ctx, FATAL, msg, meta)
}
