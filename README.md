# uuid

`uuid` is a go package that provides generation and typing for UUIDs. Sub-packages contain implementations for creating different types of UUIDs.

## Usage

```go
import (
  "github.com/leighmcculloch/go-uuid"
  "github.com/leighmcculloch/go-uuid/timeuuid"
)
```


```go
id := timeuuid.Now()
s := id.String() // e.g. 08827178-0ad4-11e7-b5df-b3f54921aa61
```

## Development

Run tests and benchmarks with:

```
make
```
