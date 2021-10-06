# perr
Perr gives error persona and enrich error.<br/>
You can handle error properly, safely and easily with perr.

## Purpose
I make this library for the following reasons.

### main purpose

- I think errors must possess their persona.
  - For client(which means client side and user using your service)
  - For developer

For example, `dial tcp 127.0.0.1:3306: connect: connection refused` must not be shown for client. At the same time, `Internal Server Error` is not enough for developer.

### sub purpose

- `Map()` and `Json()` method make it easy to store and analyze error for you.<br/>
- `Is()` and `Level()` method make it easy to handle error.<br/>
- You can trace error with `Traces()` method.

## How to use

I'll show you how to use simply.

### wrap error
```go:wrap.go
package main

import (
	"fmt"
	"strconv"
	
	"github.com/maru44/perr"
)

func main() {
	_, err := strconv.Atoi("a")

	p := perr.Wrap(err, perr.BadRequest)
	p2 := perr.Wrap(p, perr.BadRequest)
	p3 := perr.Wrap(p2, perr.InternalServerError)

	fmt.Println("developer:", p3)
	fmt.Println("client:", p3.Output())
	
	fmt.Println("\n", p3.Traces())
}

/* output */
// developer: strconv.Atoi: parsing "a": invalid syntax
// client: Bad Request

// /tmp/sandbox2199832404/prog.go:13 ===> main
//	/tmp/sandbox2199832404/prog.go:14 ===> main
//		/tmp/sandbox2199832404/prog.go:15 ===> main

```

### new error
```go:new.go
package main

import (
	"fmt"
	
	"github.com/maru44/perr"
)

func main() {
	p:= perr.New("Someone pour coffee", perr.IAmATeaPot)
	
	fmt.Println("developer: ", p)
	fmt.Println("client: ", p.Output())
	
	p2 := perr.New("", perr.IAmATeaPot)
	fmt.Println("developer: ", p2)
}

```

Here's more sample.
https://github.com/maru44/perr/tree/master/samples

