package main

import (
  "app/logger"
)

func main() {
  errSeverity := 0
  infoSeverity := 1

  l := logger.Logger{
    Label: "GCP simple example",
    ThresholdSeverity: infoSeverity,
    Severities: map[int]string{
      errSeverity: "ERROR",
      infoSeverity: "INFO",
    },
    Serialize: logger.GCPSerialize,
  }

  l.LogWithSeverity(map[string]string{"message": "Hello World"}, infoSeverity)
}


