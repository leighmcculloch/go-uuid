# timeuuid

`timeuuid` is a go package that provides generation and typing for a version 1 UUID, that uses a random non-locking clock sequence, and a random node ID.

## Usage

### Raw simple ID

```
id := timeuuid.Now()
s := id.String() // e.g. 08827178-0ad4-11e7-b5df-b3f54921aa61
```

## Development

Run tests and benchmarks with:

```
make
```
