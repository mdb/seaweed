# seaweed

A thin, Golang-based [Magic Seaweed API](http://magicseaweed.com/developer/forecast-api) client for fetching marine forecast data.

# Usage

```
go get github.com/mdb/seaweed
```

```
import (
  "github.com/mdb/seaweed"
)

func main() {
  client := seaweed.NewClient("<YOUR_API_KEY>")
  resp, err := client.Tomorrow("<SOME_SPOT_ID>")
  if err != nil {
    fmt.Println(err)
  }
}
```
