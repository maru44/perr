# perr
Perr gives error persona and enrich error.
You can handle error properly, safely and easily with perr.

## How to use perr.

Here's sample.
https://github.com/maru44/perr/tree/master/samples

```go
// wrap error
_, err := strconv.Atoi("a")
if err != nil {
    p := perr.Wrap(err, perr.FlagBadRequest)
}

// new error
p := perr.New("new error", perr.FlagIAmTeaPot)
```
