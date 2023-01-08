<h1 align="center">Go Polybase Client</h1>

## Setup

```bash
go get github.com/durudex/go-polybase@latest
```

## Usage

```go
import (
    "context"

    "github.com/durudex/go-polybase"
)

type Model struct{}

func main() {
	db := polybase.New(polybase.Config{
        URL: polybase.TestnetURL,
	})
    coll := db.Collection("Collection")

    var response polybase.SingleResponse[Model]

    coll.Record("id").Get(context.Background(), &response)

    ...
}
```

> More usage examples can be found in [examples](examples/).

## ⚠️ License

Copyright © 2022-2023 [Durudex](https://github.com/durudex). Released under the MIT license.
