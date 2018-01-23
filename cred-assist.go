package main

import "github.com/jrstill/cred-assist/data"
import "fmt"

func main() {
	// a few dummy lines so the compiler doesn't complain about anything,
	//but also so I know everything at least compiles and does "something"
	a := data.Account{
		Name:      "test",
		AccessKey: "testacc",
	}

	a.AccessKey = "1234"

	fmt.Println(a.AccessKey)
}
