package main

import (
  logger "app"
)

type Error struct {
  msg string
  trace string
}

func (e *Error) Msg() string {
  return e.msg
}

func (e *Error) Trace() string {
  return e.trace
}

func main() {
  errLogLevel := 0
  infoLogLevel := 1

  l := logger.Logger{
    WhoAmI: "example",
    Severity: infoLogLevel,
    Severities: map[int]string{
      errLogLevel: "ERROR",
      infoLogLevel: "INFO",
    },
  }

  l.LogWithSeverity("Hello World", infoLogLevel)
  l.LogError(&Error{msg: "I am an error", trace: "I am a trace"}, errLogLevel)
}
