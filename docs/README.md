# `go-polybase`

## Setup

To get the [`go-polybase`](https://github.com/durudex/go-polybase) module, you need to have or install Go version >= [1.18](https://go.dev/dl/). To check your current version of Go, use the `go version` command.

**The command to get the module:**

```bash
go get github.com/durudex/go-polybase@latest
```

## Basic Example

An example in which a new instance of the client is created and the values of the collection are obtained.

```go
import (
    "context"

    "github.com/durudex/go-polybase"
)

type Model struct { ... }

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

> More usage examples can be found in the [examples directory](../examples/README.md).

## PolyGen

For easy and fast development, you can generate code to interact with Polybase collections
using [PolyGen](https://github.com/durudex/polygen).

## License

Copyright © 2022-2023 [Durudex](https://github.com/durudex). Released under the MIT license.
