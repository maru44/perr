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

Here's sample.
https://github.com/maru44/perr/tree/master/samples

