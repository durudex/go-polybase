# `durudex/go-polybase`

Implementation of a client for the Polybase decentralized database in Go programming language.

## Setup

To get the [`go-polybase`](https://github.com/durudex/go-polybase) module, you need to have or install [Go version >= 1.19](https://go.dev/dl/). To check your current version of Go, use the `go version` command.

**The command to get the module:**

```bash
go get github.com/durudex/go-polybase@latest
```

## Example

An example where a new instance of a client, collection, record is created, and the result of the record creation is obtained.

```go
import (
    "context"

    "github.com/durudex/go-polybase"
)

// Describe the model of the collection you want to get.
type Model struct { ... }

func main() {
    // Create an instance of the client with the specified configuration.
    client := polybase.New(&polybase.Config{
        URL: polybase.TestnetURL,
    })
    // Create an instance of the collection.
    coll := polybase.NewCollection[Model](client, "Collection")

    // Create an instance of the record with the specified ID.
    record := coll.Record("id")
    // Get this record from the collection.
    response := record.Get(context.Background())

    ...
}
```

> **Note**
> More examples can be found in the [examples directory](https://github.com/durudex/go-polybase/blob/main/examples/README.md).

## License

Copyright © 2022-2023 [Durudex](https://github.com/durudex). Released under the MIT license.
