package log

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var _ Exporter = &StdExporter{}

type StdExporter struct{}

func (s StdExporter) Export(line map[string]string) {
	b, err := json.Marshal(line)
	if err != nil {
		return
	}
	fmt.Fprintln(os.Stdout, string(b))
}

// --------------FileExporter------------------- //

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)
var _ Exporter = &FileExporter{}

type FileExporter struct {
	writerCh chan []byte
}

func NewFileExporter() FileExporter {
	exp := FileExporter{
		writerCh: make(chan []byte),
	}
	go exp.loop()
	return exp
}

func (e FileExporter) Export(line map[string]string) {
	go func() {
		b, err := json.Marshal(line)
		if err != nil {
			return
		}
		e.writerCh <- b
	}()

	ts := line["timestamp"]
	delete(line, "timestamp")
	lvl := line["lvl"]
	delete(line, "lvl")
	ser := line["service.name"]
	delete(line, "service.name")
	msg := line["msg"]
	delete(line, "msg")

	fmt.Fprintln(os.Stdout, fmt.Sprintf("%s: [%s] [%s] '%s' args: %v", ts, lvl, ser, msg, line))
}

func (e FileExporter) loop() {
	for {
		select {
		case b := <-e.writerCh:
			logFilePath := strings.Replace(basepath, "pkg/log", "log.txt", 1)
			f, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				fmt.Println(err)
				return
			}

			if _, err = f.WriteString(string(b) + "\n"); err != nil {
				panic(err)
			}

			if err := f.Close(); err != nil {
				fmt.Fprintln(os.Stdout, fmt.Sprintf("unable to close file - %s", err.Error()))
			}
		}
	}
}
