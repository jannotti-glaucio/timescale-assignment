package excepts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThrowException(t *testing.T) {
	t.Run("plain message", func(t *testing.T) {
		err := ThrowException("ERR-001", "error message")

		assert.Equal(t, "ERR-001", err.Code)
		assert.Equal(t, err.Error(), "error message")
	})
	t.Run("message with parameters", func(t *testing.T) {
		err := ThrowException("ERR-002", "missing field %s", "field1")

		assert.Equal(t, "ERR-002", err.Code)
		assert.Equal(t, err.Error(), fmt.Sprintf("missing field %s", "field1"))
	})
}
