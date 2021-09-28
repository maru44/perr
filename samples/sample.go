package main

import (
	"fmt"
	"strconv"

	"github.com/maru44/perr"
)

func wrapSample() *perr.Err {
	_, err := strconv.Atoi("sample")
	return perr.Wrap(err, perr.FlagBadRequest)
}

func main() {
	fmt.Println("=================== Wrap error ===================")
	sample0 := wrapSample()
	fmt.Printf("Client: %v / Developer: %v\n\n", sample0.Output().Error(), sample0.Error())

	fmt.Println("=================== New error ===================")
	sample1 := perr.New("Someone pour coffee into tea cup.", perr.FlagIAmTeaPot)
	fmt.Printf("Client: %v / Developer: %v\n\n", sample1.Output().Error(), sample1.Error())

	fmt.Println("=================== New error with custom message ===================")
	sample2 := perr.New("Someone pour coffee into tea cup.", perr.FlagIAmTeaPot, "Don't pour coffee!")
	fmt.Printf("Client: %v / Developer: %v\n\n", sample2.Output().Error(), sample2.Error())

	fmt.Println("=================== Judge whether the error cause Perror ===================")
	fmt.Println("Whether I'm a teapot?:", sample2.Is(perr.IAmATeaPot))
	fmt.Println("Whether Bad request?:", sample2.Is(perr.BadRequest))

	fmt.Println("\n=================== Level ===================")
	sample4 := perr.New("dangerous", perr.FlagInternalServerErrorWithUrgency)
	fmt.Println(sample4.Level())
	sample5 := perr.New("not dangerous but system error", perr.FlagInternalServerError)
	fmt.Println(sample5.Level())
	sample6 := perr.New("caused by client", perr.FlagBadRequest)
	fmt.Println(sample6.Level())
	fmt.Println(wrapSample().Level())

	fmt.Println("\n=================== nil ===================")
	sample7 := perr.Wrap(nil, perr.FlagBadGateway)
	fmt.Println(sample7)
}
