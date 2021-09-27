package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/maru44/perr"
)

func wrapSample() *perr.Err {
	_, err := strconv.Atoi("sample")
	return perr.Wrap(err, perr.FlagBadRequest)
}

func outputSample() *perr.Err {
	_, err := strconv.Atoi("sample")
	return perr.Wrap(err, perr.FlagBadRequest, "With Perr", "Have a nice day")
}

func main() {
	fmt.Println("=================== Wrap error ===================")
	sample1 := wrapSample()
	fmt.Printf("For Output: %v\n", sample1.Output().Error())
	fmt.Printf("For logging: %v\n", sample1.Error())
	fmt.Printf("Stack trace:\n%v\n", sample1.Traces())

	fmt.Println("=================== original output & dict ===================")
	sample2 := outputSample()
	fmt.Printf("For Output: %v\n", sample2.Output().Error())
	fmt.Printf("For logging: %v\n", sample2.Error())
	fmt.Printf("Dict:\n%v\n", sample2.ToDict())
	fmt.Printf("01File: %v\n", sample2.ToDict().Traces[0].File)
	fmt.Printf("Error: %v\n", sample2.ToDict().Error)
	fmt.Printf("Unwrap: %v\n", sample2.UnWrap())
	json_, _ := json.Marshal(sample2.ToDict())
	fmt.Printf("json: %v\n\n", string(json_))

	fmt.Println("=================== New error ===================")
	sample3 := perr.New("pouring coffee into tea cup", perr.FlagIAmTeaPot)
	fmt.Printf("For Output: %v\n", sample3.Output().Error())
	fmt.Printf("For logging: %v\n", sample3.Error())
	fmt.Printf("Stack trace:\n%v\n", sample3.Traces())

	fmt.Println("=================== Level ===================")
	sample4 := perr.New("dangerous", perr.FlagInternalServerErrorWithUrgency)
	fmt.Println(sample4.Level())
	sample5 := perr.New("not dangerous but system error", perr.FlagInternalServerError)
	fmt.Println(sample5.Level())
	sample6 := perr.New("caused by client", perr.FlagBadRequest)
	fmt.Println(sample6.Level())

	fmt.Println("=================== nil ===================")
	sample7 := perr.Wrap(nil, perr.FlagBadGateway)
	fmt.Println(sample7)
}

/*                       output                       */

// =================== Wrap error ===================
// For Output: Bad Request
// For logging: strconv.Atoi: parsing "sample": invalid syntax
// Stack trace:
// /home/secret/perr/samples/sample.go:13 ===> wrapSample
// /home/secret/perr/samples/sample.go:23 ===> main

// =================== original output & dict ===================
// For Output: With Perr.Have a nice day
// For logging: strconv.Atoi: parsing "sample": invalid syntax
// Dict:
// &{strconv.Atoi: parsing "sample": invalid syntax strconv.Atoi: parsing "sample": invalid syntax With Perr.Have a nice day EXTERNAL ERROR /home/secret/perr/samples/sample.go:18 ===> outputSample
// /home/secret/perr/samples/sample.go:29 ===> main
//  2021-09-28 08:01:50.1127957 +0900 JST m=+0.000146801}
// 01File: /home/secret/perr/samples/sample.go
// Error: strconv.Atoi: parsing "sample": invalid syntax
// Unwrap: strconv.Atoi: parsing "sample": invalid syntax
// json: {"error":{"Func":"Atoi","Num":"sample","Err":{}},"text":"strconv.Atoi: parsing \"sample\": invalid syntax","output":"With Perr.Have a nice day","level":"EXTERNAL ERROR","traces":[{"file":"/home/secret/perr/samples/sample.go","line":18,"name":"outputSample","program_counter":4877945},{"file":"/home/secret/perr/samples/sample.go","line":29,"name":"main","program_counter":4878772}],"occured_at":"2021-09-28T08:01:50.1127957+09:00"}

// =================== New error ===================
// For Output: I'm a teapot
// For logging: pouring coffee into tea cup
// Stack trace:
// /home/secret/perr/samples/sample.go:40 ===> main

// =================== Level ===================
// ALERT
// INTERNAL ERROR
// EXTERNAL ERROR
// =================== nil ===================
// <nil>
