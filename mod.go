package pkg

import (
	"fmt"
	"runtime/debug"

	errors "github.com/jybp/errors"
	gopkgin "gopkg.in/jybp/errors.v0"
)

func Stack() string {
	return string(debug.Stack())
}

func Panic() {
	panic("panic")
}

func GopkginStack() string {
	return fmt.Sprintf("%+v", gopkgin.Stack())
}

func ErrorsStack() string {
	return fmt.Sprintf("%+v", errors.Stack())
}
