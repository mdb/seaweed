[![Build Status](https://travis-ci.org/mdb/seaweed.svg?branch=master)](https://travis-ci.org/mdb/seaweed) [![Go Report Card](https://goreportcard.com/badge/github.com/mdb/seaweed)](https://goreportcard.com/report/github.com/mdb/seaweed)

# seaweed

A thin, Golang-based [Magic Seaweed API](http://magicseaweed.com/developer/forecast-api) client for fetching marine forecast data.

## Usage

Installation:

```
go get github.com/mdb/seaweed
```

Example:

```go
import (
  "github.com/mdb/seaweed"
  "github.com/tonnerre/golang-pretty"
)

func main() {
  client := seaweed.NewClient("<YOUR_API_KEY>")
  resp, err := client.Forecast("<SOME_SPOT_ID>")
  if err != nil {
    fmt.Println(err)
  }

  fmt.Printf("%# v", pretty.Formatter(resp))
}
```

Use a customized client:

```go
client := seaweed.Client{
  APIKey:     string,
  HttpClient: *http.Client,
  CacheAge:   time.Duration, // override 5m default
  CacheDir:   string, // override os.TempDir() value
  Log:        *logging.Logger, // override NewLogger(logging.INFO)
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

// Weekend forecast
resp, err := client.Weekend("<SOME_SPOT_ID>")
```

## Options

To disable response caching, set a `SW_DISABLE_CACHE` env var:

```
SW_DISABLE_CACHE=true
```
