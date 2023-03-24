[![CI](https://github.com/mdb/seaweed/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/mdb/seaweed/actions/workflows/ci.yaml) [![PkgGoDev](https://pkg.go.dev/badge/github.com/mdb/seaweed)](https://pkg.go.dev/github.com/mdb/seaweed) [![Go Report Card](https://goreportcard.com/badge/github.com/mdb/seaweed)](https://goreportcard.com/report/github.com/mdb/seaweed)

# seaweed

A thin, Go [Magic Seaweed API](http://magicseaweed.com/developer/forecast-api) client for fetching marine forecast data.

## Usage

Example:

```go
import (
  "github.com/mdb/seaweed"
)

func main() {
  client := seaweed.NewClient("<YOUR_API_KEY>")
  resp, err := client.Forecast("<SOME_SPOT_ID>")
  if err != nil {
    panic(err)
  }

  fmt.Printf("%# v", resp)
}
```

Use a customized client:

```go
client := seaweed.Client{
  BaseURL:    "https://magicseaweed.com",
  APIKey:     "YOUR_KEY",
  HTTPClient: &http.Client{},      // *http.Client
  Logger:     logrus.New(),        // *logrus.Logger
  clock:      seaweed.RealClock{}, // seaweed.Clock
}
```

Client methods:

```go
import (
  "github.com/mdb/seaweed"
)

client := seaweed.NewClient("<YOUR_API_KEY>")

// Full forecast
resp, err := client.Forecast("<SOME_SPOT_ID>")

// Today's forecast
resp, err := client.Today("<SOME_SPOT_ID>")

// Tomorrow's forecast
resp, err := client.Tomorrow("<SOME_SPOT_ID>")

// This weekend's forecast
resp, err := client.Weekend("<SOME_SPOT_ID>")
```
