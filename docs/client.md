# [Client](https://pkg.go.dev/github.com/durudex/go-polybase#Client)

The client is a crucial component of any application that utilizes the [`go-polybase`](https://github.com/durudex/go-polybase) module. This interface is responsible for processing and sending requests to the Polybase API.

## [New](https://pkg.go.dev/github.com/durudex/go-polybase#New)

To start using the [`go-polybase`](https://github.com/durudex/go-polybase) client, you need to crete a new client instance. This can be done using the internal [`New()`](https://pkg.go.dev/github.com/durudex/go-polybase#New) function, which returns a new instance with either a specified configuration or the default configuration.

To create an instance with a specified configuration, you need to pass a pointer of [`Config`](https://pkg.go.dev/github.com/durudex/go-polybase#Config) value as an argument to the [`New(...)`](https://pkg.go.dev/github.com/durudex/go-polybase#New) function. This can be useful if you want to use specific settings, for example, if you have your own configuration file.

**An example of using client configuration:**

```go
client := polybase.New(&polybase.Config{
    ...
})
```

If you want to use the default configuration, you can simply call the [`New()`](https://pkg.go.dev/github.com/durudex/go-polybase#New) function without any arguments. The client will be created with the default configuration set in the [`go-polybase`](https://github.com/durudex/go-polybase) module.

**An example of creating a client with a default configuration:**

```go
client := polybase.New()
```

## [`Config`](https://pkg.go.dev/github.com/durudex/go-polybase#Config)

Each of the configuration options listed below is a field in the internal [`Config`](https://pkg.go.dev/github.com/durudex/go-polybase#Config) structure used to configure the [`go-polybase`](https://github.com/durudex/go-polybase) client. You can set the options of these fields to configure the client according to your needs.

### [`Config.URL`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.URL)

The [`URL`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.URL) field defines a url to a node or any other Polybase API gateway to which the client will send requests. You can use pre-defined internal url values or specify your own url values.

**Internal values:**

- [`DefaultURL`](https://pkg.go.dev/github.com/durudex/go-polybase#DefaultURL) (Default)
- [`TestnetURL`](https://pkg.go.dev/github.com/durudex/go-polybase#TestnetURL)

### [`Config.Name`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.Name)

The [`Name`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.Name) field defines the client name used as the value of the `X-Polybase-Client` HTTP header in requests to the Polybase API.

Additionally, for better analysis, the prefix `"durudex/go-polybase:"` is added to each name. This allows for easier identification of the module or library from which requests are made.

### [`Config.DefaultNamespace`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.DefaultNamespace)

The [`DefaultNamespace`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.DefaultNamespace) field defines the default namespace that will be added to the collection name when creating a new instance.

### [`Config.RecoverHandler`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.RecoverHandler)

The [`RecoverHandler`](https://pkg.go.dev/github.com/durudex/go-polybase#Config.RecoverHandler) field defines the handler that will be called in case of a panic.

Panics usually occur during development and may indicate passing an incorrect type or a lack of connection to the internet or the Polybase API.
