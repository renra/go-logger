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
    Label: "GCP simple example",
    ThresholdSeverity: infoLogLevel,
    Severities: map[int]string{
      errLogLevel: "ERROR",
      infoLogLevel: "INFO",
    },
    Serialize: logger.GCPSerialize,
  }

  l.LogWithSeverity(map[string]string{"message": "Hello World"}, infoLogLevel)

  l.LogErrorWithSeverity(&Error{msg: "I am a msg", trace: "I am a trace"}, errLogLevel)
  l.LogErrorWithSeverity(&Error{msg: "I am a msg", trace: "I am a trace"}, -1) // prints with the UNKNOWN label
}


