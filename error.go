package main

import "fmt"

type errorWithCode struct {
	code    int
	message string
}

func (e *errorWithCode) Error() string {
	return fmt.Sprintf("%d: %s", e.code, e.message)
}
