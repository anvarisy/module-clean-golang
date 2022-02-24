package error

import "errors"

var (
	ErrPayloadValidation = errors.New("invalid payload")
)
