# Logger

This tool aims to provide a very simple leveled logging library. It allows you to define your own format of logging as well as your own log levels or severities. It does not aim to be good or fast at object serialization. Give it strings, it's a logger. Instead it aims to make your hands free in case you want to rewrite the performance-critical parts yourself.

## Usage

You can define your logger by a few key attributes ...

```go
import (
  "fmt"
  "time"
  "github.com/GlobalWebIndex/go-logger/logger"
)

func main() {
  errSeverity = 0
  infoSeverity = 1

  l := logger.Logger{
    Label: "My shiny app",           // can be used to distinguish logs from your app from other apps
    ThresholdSeverity: infoSeverity, // outputs everything that has this severity or lower
    Severities: map[int]string {     // adds severity labels which you can print, defaults to "UNKNOWN"
      errSeverity: "ERROR",
      infoSeverity: "INFO"
    },
    Serialize: func(logger * logger.Logger, data map[string]string, severity int) {
      t := time.Now().Unix()
      output := fmt.Sprintf("[%s] (%s) %d:", logger.Label, logger.SeverityLabel(severity), t)

      fmt.Println(output)
    }
  }

  l.LogWithSeverity(map[string]string{"message": "Hello World"}, infoSeverity)
}
```

This outputs: `[My shiny app] (INFO) 1549635939: message=Hello World`.

This should give you almost total flexibility to output the data the way you want. Of course it's a bit too verbose, but you can always build on top of that.

Sending a map to a logging library might seem like an overkill compared to sending just a string, but it's a more generic case that you want to log several pieces of information side-by-side - for example when logging an error message together with the stacktrace.

Logger ships with some basic functionality to output logs in a GCP-friendly manner ...

```go
l := logger.Logger{
  Label: "GCP simple example",
  ThresholdSeverity: infoSeverity,
  Severities: map[int]string{
    errSeverity: "ERROR",
    infoSeverity: "INFO",
  },
  Serialize: logger.GCPSerialize,
}

l.LogWithSeverity(map[string]string{"message": "Hello World"}, infoLogLevel)
```

In which case the output looks like this: `{"date":1549636260,"message":"Hello World","serviceContext":{"service":"GCP simple example"},"severity":"INFO"}`
