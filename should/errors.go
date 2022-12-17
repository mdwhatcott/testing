package should

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
)

var (
	ErrExpectedCountInvalid = errors.New("expected count invalid")
	ErrTypeMismatch         = errors.New("type mismatch")
	ErrKindMismatch         = errors.New("kind mismatch")
	ErrAssertionFailure     = errors.New("assertion failure")
)

func failure(format string, args ...any) error {
	return wrap(ErrAssertionFailure,
		format+"\nStack: (filtered)\n%s",
		append(args, stack())...)
}
func wrap(inner error, format string, args ...any) error {
	return fmt.Errorf("%w: "+fmt.Sprintf(format, args...), inner)
}
func stack() string {
	lines := strings.Split(string(debug.Stack()), "\n")
	var filtered []string
	for x := 1; x < len(lines)-1; x += 2 {
		if strings.Contains(lines[x+1], "_test.go:") {
			filtered = append(filtered, lines[x], lines[x+1])
		}
	}
	return "> " + strings.Join(filtered, "\n> ")
}
