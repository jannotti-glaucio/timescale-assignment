package exceptions

import (
	"errors"
	"fmt"
)

func ThrowError(message string, args ...interface{}) error {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	return errors.New(message)
}
