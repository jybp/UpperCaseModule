package main

import (
	"flag"
	"fmt"

	mod "github.com/jybp/UpperCaseModule"
)

var ldflags, stack, panic bool

func main() {
	flag.BoolVar(&stack, "stack", false, "")
	flag.BoolVar(&panic, "panic", false, "")
	flag.Parse()
	run()
}

func run() {
	if stack {
		fmt.Printf("debug.Stack:\n%s\n", mod.Stack())
		fmt.Printf("errors Stack:\n%s\n", mod.ErrorsStack())
		fmt.Printf("gopkg.in Stack:\n%s\n", mod.GopkginStack())
	}

	// No stack trace attached
	if panic {
		fmt.Printf("pkg panic:\n%+v\n", recoverPanic(mod.Panic))
	}
}

func recoverPanic(f func()) (r interface{}) {
	defer func() { r = recover() }()
	f()
	return
}
