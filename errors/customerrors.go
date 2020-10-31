package errors

import "errors"

var (
	ErrNoSuchElement = errors.New("priority queue underflow")
)
