package main

import (
  logger "app"
)

type Error struct {
  trace string
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
  l.LogErrorWithSeverity(&Error{trace: "I am a trace"}, errLogLevel)
  l.LogErrorWithSeverity(&Error{trace: "I am a trace"}, -1) // prints with the UNKNOWN label
}
