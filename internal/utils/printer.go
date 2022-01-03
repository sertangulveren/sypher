package utils

import (
	"fmt"
	"os"
	"runtime/debug"
)

func ExitWithMessage(err error, msg string) {
	if err == nil {
		return
	}
	if os.Getenv("DEBUG") != "" {
		debug.PrintStack()
	}
	fmt.Println(err)

	fmt.Println(msg)

	os.Exit(1)
}
func PanicWithError(err error) {
	if err != nil {
		panic(err)
	}
}
