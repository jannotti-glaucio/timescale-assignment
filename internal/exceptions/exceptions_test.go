package exceptions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThrowError(t *testing.T) {
	t.Run("plain message", func(t *testing.T) {
		err := ThrowError("error message")

		assert.Equal(t, err.Error(), "error message")

	})
	t.Run("message with parameters", func(t *testing.T) {
		err := ThrowError("missing field %s", "field1")

		assert.Equal(t, err.Error(), fmt.Sprintf("missing field %s", "field1"))
	})
}
