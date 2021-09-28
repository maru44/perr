# perr
Perr gives error persona and enrich error.<br/>
You can handle error properly, safely and easily with perr.

## purpose
I make this library for the following reasons.

### main purpose

- I think errors must possess their persona.
  - For client(which means client side and user using your service)
  - For developer

For example, `dial tcp 127.0.0.1:3306: connect: connection refused` must not be shown for client. At the same time, `Internal Server Error` is not enough for developer.

### sub purpose

- `Map()` and `Json()` method make it easy to store and analyze error for you.<br/>
- `Is()` method and `Level` make it easy to handle error.<br/>
- You can trace error with `Trace()` method.

## How to use perr.

Here's sample.
https://github.com/maru44/perr/tree/master/samples

```go:sample.go
package main

import (
	"fmt"
	"strconv"

	"github.com/maru44/perr"
)

func main() {
	_, err := strconv.Atoi("sample")

	/* =================== Wrap error =================== */
	fmt.Println("=================== Wrap error ===================")
	sample0 := perr.Wrap(err, perr.BadRequest)
	fmt.Printf("Client: %v / Developer: %v\n\n", sample0.Output().Error(), sample0.Error())

	// output >>
	// Client: Bad Request / Developer: strconv.Atoi: parsing "sample": invalid syntax

	/* =================== New error =================== */
	fmt.Println("=================== New error ===================")
	sample1 := perr.New("Someone pour coffee into tea cup.", perr.IAmATeaPot)
	fmt.Printf("Client: %v / Developer: %v\n\n", sample1.Output().Error(), sample1.Error())

	// output >>
	// Client: I'm a teapot / Developer: Someone pour coffee into tea cup.

	/* =================== New error with custom message =================== */
	fmt.Println("=================== New error with custom message ===================")
	sample2 := perr.New("Someone pour coffee into tea cup.", perr.IAmATeaPot, "Don't pour coffee!")
	fmt.Printf("Client: %v / Developer: %v\n\n", sample2.Output().Error(), sample2.Error())

	// output >>
	// Client: Don't pour coffee! / Developer: Someone pour coffee into tea cup.

	/* =================== Judge whether the error cause Perror =================== */
	fmt.Println("=================== Judge whether the error cause Perror ===================")
	fmt.Println("Whether I'm a teapot?:", sample2.Is(perr.IAmATeaPot))
	fmt.Println("Whether Bad request?:", sample2.Is(perr.BadRequest))

	// output >>
	// Whether I'm a teapot?: true
	// Whether Bad request?: false

	/* =================== Level =================== */
	fmt.Println("\n=================== Level ===================")
	sample4 := perr.New("dangerous", perr.InternalServerErrorWithUrgency)
	fmt.Println(sample4.Level)
	sample5 := perr.New("not dangerous but system error", perr.InternalServerError)
	fmt.Println(sample5.Level)
	sample6 := perr.New("caused by client", perr.BadRequest)
	fmt.Println(sample6.Level)
	fmt.Println(sample0.Level)

	// output >>
	// ALERT
	// INTERNAL ERROR
	// EXTERNAL ERROR
	// EXTERNAL ERROR

	/* =================== with level =================== */
	fmt.Println("\n=================== with level ===================")
	myLevel := perr.ErrLevel("My Level")
	sample7 := perr.NewWithLevel("this is my level", perr.IAmATeaPot, myLevel)
	fmt.Println(sample7.Level)

	// output >>
	// My Level

	/* =================== wrap nil =================== */
	fmt.Println("\n=================== wrap nil ===================")
	sample8 := perr.Wrap(nil, perr.BadGateway)
	fmt.Println(sample8)

	// output >>
	// <nil>

	/* =================== Stack trace =================== */
	fmt.Println("\n=================== Stack trace ===================")
	sample9 := perr.Wrap(err, perr.BadRequest)
	fmt.Printf("stacktrace:\n%v\n", sample9.Traces())

	// output >>
	// stacktrace:
	// /secret/perr/samples/sample.go:81 ===> main

	/* =================== Map & Json =================== */
	fmt.Println("\n=================== Map & Json ===================")
	fmt.Printf("map:\n%v\n", sample2.Map())
	json_ := sample2.Json()
	fmt.Printf("json:\n%v\n\n", string(json_))

	// output >>
	// 	map:
	// &{I'm a teapot Don't pour coffee! Someone pour coffee into tea cup. Don't pour coffee! EXTERNAL ERROR /secret/perr/samples/sample.go:31 ===> main
	//  2021-09-28 21:37:06.7194034 +0900 JST m=+0.000158501}
	// json:
	// {"error":"I'm a teapot","teated_as":"Don't pour coffee!","msg_for_developer":"Someone pour coffee into tea cup.","msg_for_client":"Don't pour coffee!","level":"EXTERNAL ERROR","traces":[{"file":"/secret/perr/samples/sample.go","line":31,"name":"main","program_counter":4886631}],"occured_at":"2021-09-28T21:37:06.7194034+09:00"}
}

```
