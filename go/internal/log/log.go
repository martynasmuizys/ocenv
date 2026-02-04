package log

import (
	"fmt"
	"os"
)

func Hprint(args ...any) {
	fmt.Printf(":: ")
	fmt.Println(args...)
}

func Printf(format string, args ...any) {
	fmt.Printf("==> "+format, args...)
}

func Println(args ...any) {
	fmt.Printf("==> ")
	fmt.Println(args...)
}

func Fatal(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
	os.Exit(1)
}
