# Вхідні дані

Вхідні дані представляють аргументи функції у колекції Polybase.

## Використання

Щоб розпочати використання, вам потрібно знати, які аргументи приймає функція, яку ви плануєте викликати. Розглянемо приклад функції-конструктора:

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

Отже, вам потрібно передати такі аргументи: `id`, `name`, `age`. Це можна зробити декількома способами.

### Значення

```go
var (
    id = "1"
    name = "example"
    age = 146
)

coll.Create(..., args: id, name, status)
```

### Структура

```go
type Input struct {
    ID   string
    Name string
    Age  int
}

input := Input{"1", "example", 146}

coll.Create(..., args: input)
```

### Вказівник

```go
var (
    id = "1"
    name = "example"
    age = 146
)

coll.Create(..., args: &id, &name, &age)
```

### Масив

```go
input := []any{"1", "example", 146}
coll.Create(..., args: input)
```

Зверніть увагу, що масив або зріз повинен мати тип `any`, інакше цей масив або зріз буде використано як один аргумент.

## Приклад

Цей приклад призначений виключно для демонстрації можливостей аналізатора аргументів.

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
        balance: map<string, number>
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

**Кінцевий результат, який буде надіслано до API Polybase, матиме наступний вигляд:**

```json
["uid:1", 146, true, ["durudex", "blockchain"], {"DUR": 1}]
```
