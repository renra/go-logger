package main

import (
  "fmt"
  "time"
  "strings"
  "app/logger"
)

func main() {
  infoSeverity := 0

  l := logger.Logger{
    Label: "simple example",
    ThresholdSeverity: infoSeverity,
    Severities: map[int]string{
      infoSeverity: "INFO",
    },
    Serialize: func(logger *logger.Logger, data map[string]string, severity int) {
      t := time.Now().Unix()
      output := fmt.Sprintf("[%s] (%s) %d:", logger.Label, logger.SeverityLabel(severity), t)

      for k, v := range data {
        output = strings.Join([]string{output, fmt.Sprintf("%s=%s", k, v)}, " ")
      }

      fmt.Println(output)
    },
  }

  l.LogWithSeverity(map[string]string{"message": "Hello World"}, infoSeverity)
}

