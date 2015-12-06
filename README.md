[![Build Status](https://travis-ci.org/mdb/seaweed.svg?branch=master)](https://travis-ci.org/mdb/seaweed)

# seaweed

A thin, Golang-based [Magic Seaweed API](http://magicseaweed.com/developer/forecast-api) client for fetching marine forecast data.

## Usage

Installation:

```
go get github.com/mdb/seaweed
```

Example:

```
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

Client methods:

```
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
```

## Options

To log request/response details, set a `SW_LOG` env var:

```
SW_LOG=true
```

To disable 5 minute response caching, set a `SW_DISABLE_CACHE` env var:

```
SW_DISABLE_CACHE=true
```
