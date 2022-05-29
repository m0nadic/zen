package util

import (
	"fmt"
	"os"
)

func IfErrorExit(err error, code int) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(code)
	}
}
