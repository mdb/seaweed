# seaweed

A thin, Golang-based [Magic Seaweed API](http://magicseaweed.com/developer/forecast-api) client for fetching marine forecast data.

# Usage

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
