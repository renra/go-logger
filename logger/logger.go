package logger

import (
  "fmt"
  "time"
  "encoding/json"
)

const UnknownLabel = "UNKNOWN"

type Logger struct {
  Label string
  ThresholdSeverity int
  Severities map[int]string
  Serialize func(*Logger, map[string]string, int)
}

func (l *Logger) LogWithSeverity(data map[string]string, severity int) {
  if l.ThresholdSeverity >= severity {
    l.Serialize(l, data, severity)
  }
}

func (l *Logger) MergeWithBasicDataForGCP(data map[string]string, severity int) map[string]interface{} {
  output := map[string]interface{}{
    "date": time.Now().Unix(),
    "serviceContext": map[string]string{"service": l.Label},
    "severity": l.SeverityLabel(severity),
  }

  for k, v := range data {
    output[k] = v
  }

  return output
}

func GCPSerialize(l *Logger, data map[string]string, severity int) {
  d := l.MergeWithBasicDataForGCP(data, severity)
  bytes, _ := json.Marshal(d)

  fmt.Println(string(bytes))
}

func (l *Logger) SeverityLabel(severity int) string {
  severityLabel, found := l.Severities[severity]

  if found == false {
    return UnknownLabel
  } else {
    return severityLabel
  }
}
