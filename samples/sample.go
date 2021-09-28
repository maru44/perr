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

	/* =================== New error =================== */
	fmt.Println("\n=================== nil ===================")
	sample7 := perr.Wrap(nil, perr.BadGateway)
	fmt.Println(sample7)

	// output >>
	// <nil>

	/* =================== New error =================== */

}
