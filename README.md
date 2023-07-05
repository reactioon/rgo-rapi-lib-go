# RAPI Library / Go

This is the official repository of the Reactioon API communication library for the Go language.

## Focus
This library must meet some requirements, see the list below:

1. Simple and easy to use.
2. Easy to maintain/upgrade.
3. Reusable

## Usage
You can use the library with two ways, doing an method chaining to an specific context or using method reference to create your context using the same keys.

1. Method chaining

```go
import "github.com/reactioon/rgo-rapi-lib-go/rapi"
requestReturnGet, requestErrGet := rapi.Load([]byte("your-key"), []byte("your-secret")).Request("GET", "api/v2/bots/spot/all", make(map[string]string))
```

2. Method Reference

```go
import "github.com/reactioon/rgo-rapi-lib-go/rapi"
r := rapi.Load([]byte("your-key"), []byte("your-secret"))
requestReturnGet, requestErrGet := r.Request("GET", "api/v2/bots/spot/all", make(map[string]string))
```

## Considerations
This library is under development and may change over time. The integrity of existing methods will be maintained to avoid compatibility issues in the future.

## Contributions
You can contribute to the development of the ecosystem by helping to improve this library. Feel free to improve and submit your work with a pull request.


@author José Wilker <josewilker@reactioon.com>