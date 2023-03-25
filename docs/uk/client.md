# [Client](https://pkg.go.dev/github.com/durudex/go-polybase#Client)

Клієнт є ключовою складовою будь-якого додатка, що використовує модуль [`go-polybase`](https://github.com/durudex/go-polybase). Цей інтерфейс відповідає за обробку та передачу запитів до API Polybase.

## [New](https://pkg.go.dev/github.com/durudex/go-polybase#New)

Для того щоб розпочати використання клієнта [`go-polybase`](https://github.com/durudex/go-polybase), вам необхідно створити новий екземпляр клієнта. Це можна зробити за допомогою внутрішньої функції [`New()`](https://pkg.go.dev/github.com/durudex/go-polybase#New), яка поверне новий екземпляр з вказаною конфігурацією або конфігурацією замовчуванням.

Для створення екземпляру з вказаною конфігурацією, вам потрібно передати вказівник на значення [`Config`](https://pkg.go.dev/github.com/durudex/go-polybase#Config) в якості аргументу функції [`New(...)`](https://pkg.go.dev/github.com/durudex/go-polybase#New). Це може бути корисно, якщо ви хочете використовувати певні налаштування, наприклад, якщо у вас є власний файл конфігурації.

**Приклад використання конфігурації клієнта:**

```go
client := polybase.New(&polybase.Config{
    ...
})
```

Якщо ви хочете використовувати конфігурацію за замовчуванням, ви можете просто викликати функцію [`New()`](https://pkg.go.dev/github.com/durudex/go-polybase#New) без аргументів. Клієнт буде створено з конфігурацією за замовчуванням, яка встановлена у модулі [`go-polybase`](https://github.com/durudex/go-polybase).

**Приклад створення клієнта з конфігурацією за замовчуванням:**

```go
client := polybase.New()
```

## [`Config`](https://pkg.go.dev/github.com/durudex/go-polybase#Config)

Кожне з перелічених нижче значень конфігурації є полями внутрішньої структури [`Config`](https://pkg.go.dev/github.com/durudex/go-polybase#Config), яка використовується для налаштування клієнта [`go-polybase`](https://github.com/durudex/go-polybase). Ви можете встановлювати значення цих полів, щоб налаштувати клієнт згідно зі своїми потребами.

### [`Config.URL`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.URL)

Поле [`URL`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.URL) визначає посилання на вузол або будь-який інший шлюз API Polybase, до якого клієнт буде надсилати запити. Ви можете використовувати готові внутрішні значення посилання або вказати свої власні значення.

**Внутрішні значення:**

- [`DefaultURL`](https://pkg.go.dev/github.com/durudex/go-polybase#DefaultURL) (За замовчуванням)
- [`TestnetURL`](https://pkg.go.dev/github.com/durudex/go-polybase#TestnetURL)

### [`Config.Name`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.Name)

Поле [`Name`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.Name) визначає імя клієнта, яке використовується як значення HTTP-заголовка `X-Polybase-Client` для запитів до API Polybase.

Додатково для покращення аналізу, до кожного імені додається префікс `"durudex/go-polybase:"`. Це дозволяє легше визначати модуль або бібліотеку з якої здійснюються запити.

### [`Config.DefaultNamespace`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.DefaultNamespace)

Поле [`DefaultNamespace`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.DefaultNamespace) визначає простір імен за замовчуванням, який буде додаватися до імені колекції при створенні нового екземпляру.

### [`Config.RecoverHandler`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.RecoverHandler)

Поле [`RecoverHandler`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.RecoverHandler) визначає обробник, який буде викликаний у випадку виникнення паніки.

Паніки зазвичай виникають під час розробки та можуть свідчити про передачу хибного типу або відсутність з'єднання з Інтернетом або API Polybase.
