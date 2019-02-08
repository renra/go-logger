package logger

import (
  "fmt"
  "time"
)

const UnknownLabel = "UNKNOWN"

type Logger struct {
  WhoAmI string
  Severity int
  Severities map[int]string
}

const GCPTemplate = "{\"severity\":\"%s\",\"date\":%d,\"serviceContext\":{\"service\":\"%s\"},\"message\":\"%s\"}"

func (l *Logger) LogWithSeverity(msg string, severity int) {
  if l.Severity >= severity {
    severityLabel, found := l.Severities[severity]

    if found == false {
      severityLabel = UnknownLabel
    }

    timestamp := time.Now().Unix()

    fmt.Println(fmt.Sprintf(
      GCPTemplate,
      severityLabel, timestamp, l.WhoAmI,
      msg,
    ))
  }
}

type Error interface {
  Trace() string
}

func (l *Logger) LogErrorWithSeverity(e Error, severity int) {
  msg := fmt.Sprintf("{\"trace\":\"%s\"}", e.Trace())
  l.LogWithSeverity(msg, severity)
}

