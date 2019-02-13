package main

import (
  "fmt"
  "encoding/json"
  "app/logger"
)

func main() {
  errSeverity := 0
  infoSeverity := 1

  l := logger.Logger{
    Label: "GCP example",
    ThresholdSeverity: infoSeverity,
    Severities: map[int]string{
      errSeverity: "ERROR",
      infoSeverity: "INFO",
    },
    Serialize: func(logger *logger.Logger, data map[string]string, severity int) {
      d := logger.MergeWithBasicDataForGCP(data, severity)
      bytes, _ := json.Marshal(d)

      fmt.Println(string(bytes))
    },
  }

  l.LogWithSeverity(map[string]string{"message": "Hello World"}, infoSeverity)
}

