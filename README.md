# publicid

URL-safe non-sequential public ID generator using a fixed `[0-9a-z]` nanoid alphabet and a default length of 24.

Sized for collision-free use at scale per [the PlanetScale write-up](https://planetscale.com/blog/why-we-chose-nanoids-for-planetscales-api).

## Install

```sh
go get github.com/ubgo/publicid
```

## Quick start

```go
package main

import (
    "fmt"

    "github.com/ubgo/publicid"
)

func main() {
    id, err := publicid.New()
    if err != nil { panic(err) }
    fmt.Println(id) // e.g. "k7n3p9q2x8m4r6t1v5y0w8z3"
}
```

For places where the error return is awkward (struct literals, default fields):

```go
fmt.Println(publicid.Must())
```

## Custom length

```go
id, err := publicid.NewN(12)        // 12-char id
short := publicid.MustN(7)          // 7-char id, panics on bad length
```

## Validation

`Validate` checks that the value is a non-empty string of the default length, made up only of characters from `Alphabet`. The `fieldName` parameter appears in error messages so debugging is easy.

```go
if err := publicid.Validate("user_id", input); err != nil {
    return err
}
// or with custom length:
if err := publicid.ValidateN("short_code", code, 7); err != nil { ... }
```

## API

| Symbol | Purpose |
|--------|---------|
| `Alphabet` | The character set used by every public ID. |
| `DefaultLength` | The default length (24) used by `New` / `Must` / `Validate`. |
| `New() (string, error)` | Generate an ID of default length. |
| `Must() string` | Same as `New` but panics on error. |
| `NewN(n int) (string, error)` | Generate an ID of length `n`. |
| `MustN(n int) string` | Same as `NewN` but panics on error or `n <= 0`. |
| `Validate(fieldName, id string) error` | Check that `id` is non-empty, default length, and uses only `Alphabet`. |
| `ValidateN(fieldName, id string, n int) error` | Same as `Validate` but checks for length `n`. |

## License

Apache License 2.0. See [`LICENSE`](./LICENSE).
