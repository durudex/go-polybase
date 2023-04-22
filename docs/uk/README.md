# `durudex/go-polybase`

Реалізація клієнта для децентралізованої бази даних Polybase на мові Go.

## Встановлення

Щоб отримати модуль [`go-polybase`](https://github.com/durudex/go-polybase), вам необхідно мати або встановити версію [Go >= 1.19](https://go.dev/dl/). Щоб перевірити поточну версію Go, використовуйте команду `go version`.

**Команда для встановлення модуля:**

```bash
go get github.com/durudex/go-polybase@latest
```

## Приклад

Приклад, в якому створюється новий екземпляр клієнта, колекції, записа та отримується результат запису.

```go
import (
    "context"

    "github.com/durudex/go-polybase"
)

// Опишіть модель колекції, яку ви хочете отримати.
type Model struct { ... }

func main() {
    // Створіть екземпляр клієнта з вказаною конфігурацією.
    client := polybase.New(&polybase.Config{
        URL: polybase.TestnetURL,
    })
    // Створюємо екземпляр колекції.
    coll := polybase.NewCollection[Model](client, "Collection")

    // Створюємо екземпляр записі з вказаним ідентифікатором.
    record := coll.Record("id")
    // Отримуємо цей запис з колекції.
    response := record.Get(context.Background())

    ...
}
```

> **Note**
> Додаткові приклади можна знайти у [каталозі examples](https://github.com/durudex/go-polybase/blob/main/examples/README.md).

## Ліцензія

Авторське право © 2022-2023 [Durudex](https://github.com/durudex). Випущено за ліцензією MIT.
