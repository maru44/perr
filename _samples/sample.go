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
	fmt.Println(sample4.Level())
	sample5 := perr.New("not dangerous but system error", perr.InternalServerError)
	fmt.Println(sample5.Level())
	sample6 := perr.New("caused by client", perr.BadRequest)
	fmt.Println(sample6.Level())
	fmt.Println(sample0.Level())

	// output >>
	// ALERT
	// INTERNAL ERROR
	// EXTERNAL ERROR
	// EXTERNAL ERROR

	/* =================== with level =================== */
	fmt.Println("\n=================== with level ===================")
	myLevel := perr.ErrLevel("My Level")
	sample7 := perr.NewWithLevel("this is my level", perr.IAmATeaPot, myLevel)
	fmt.Println(sample7.Level())

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
	fmt.Printf("stacktrace:\n%v\n", sample9.Traces().String())

	// output >>
	// stacktrace:
	// /secret/perr/samples/sample.go:81 ===> main

	/* =================== Map & Json =================== */
	fmt.Println("\n=================== Map & Json ===================")
	fmt.Printf("map:\n%v\n", sample2.Map())
	json_ := sample2.Json()
	fmt.Printf("json:\n%v\n\n", string(json_))

	// output >>
	// map:
	// &{Someone pour coffee into tea cup. I'm a teapot Don't pour coffee! EXTERNAL ERROR /secret/perr/samples/sample.go:31 ===> main
	//  2021-09-29 18:52:04.1909594 +0900 JST m=+0.000143801}
	// json:
	// {"error":"Someone pour coffee into tea cup.","treated_as":"I'm a teapot","msg_for_client":"Don't pour coffee!","level":"EXTERNAL ERROR","traces":[{"file":"/secret/perr/samples/sample.go","line":31,"name":"main","program_counter":4885497}],"occured_at":"2021-09-29T18:52:04.1909594+09:00"}

	/* =================== judge is Perror =================== */
	fmt.Println(perr.IsPerror(err))
	p0, ok := perr.IsPerror(sample0)

	fmt.Println(ok)
	fmt.Println(string(p0.Json()))

	// output >>
	// <nil> false
	// true
	// {"error":"strconv.Atoi: parsing \"sample\": invalid syntax","treated_as":"Bad Request","msg_for_client":"Bad Request","level":"EXTERNAL ERROR","traces":[{"file":"/home/takumaru/codes/perr/samples/sample.go","line":15,"name":"main","program_counter":4880005,"layer":0}],"occurred_at":"2021-09-30T21:22:47.9695069+09:00"}

	over1 := perr.Wrap(sample0, nil)
	fmt.Println(over1.Traces())
	fmt.Printf("%d ---> Layer is %d", over1.Traces()[0].Line, over1.Traces()[0].Layer)
	fmt.Printf("\n%d ---> Layer is %d\n", over1.Traces()[1].Line, over1.Traces()[1].Layer)

	// output >>
	// /home/takumaru/codes/perr/samples/sample.go:15 ===> main
	//        /home/takumaru/codes/perr/samples/sample.go:110 ===> main
	//
	// 15 ---> Layer is 0
	// 110 ---> Layer is 1
}
