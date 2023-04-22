# Input

Input data represents the arguments of the function in the Polybase collection.

## Usage

To get started, you need to find out what arguments the function you want to call takes. Let's consider an example with a constructor function:

```polylang
collection Example {
    id: string;
    ...

    constructor(
        id: string,
        name: string,
        age: number,
    ) { ... }
}
```

So, you need to pass the following arguments: `id`, `name`, `age`. This can be done in several ways.

### Variable

```go
var (
    id = "1"
    name = "example"
    age = 146
)

coll.Create(..., args: id, name, status)
```

### Structure

```go
type Input struct {
    ID   string
    Name string
    Age  int
}

input := Input{"1", "example", 146}
coll.Create(..., args: input)
```

### Pointer

```go
var (
    id = "1"
    name = "example"
    age = 146
)

coll.Create(..., args: &id, &name, &age)
```

### Array

```go
input := []any{"1", "example", 146}
coll.Create(..., args: input)
```

> **Note**
> Array or slice must be of type `any`, otherwise this array or slice will be used as a single argument.

## Example

This example is intended solely to demonstrate the capabilities of the argument parser.

```polylang
collection Example {
    id: string;
    info: {
        age: number;
        alive: boolean;
    },
    tags: string[];
    balance: map<string, number>;

    constructor(
        id string,
        age number,
        alive boolean,
        tags string[],
        balance: map<string, number>,
    ) { ... }
}
```

```go
type Info struct {
    Age   int
    Alive bool
}

var (
    id = "uid:1"
    info = Info{146, true}
    tags = []string{"durudex", "blockchain"}
    balance = map[string]int{"DUR": 1}
)

coll.Create(..., args: id, info, tags, balance)
```

**The final result that will be sent to the Polybase API will have the following format:**

```json
["uid:1", 146, true, ["durudex", "blockchain"], {"DUR": 1}]
```
