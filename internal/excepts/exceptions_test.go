//go:build !integration

package excepts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThrowException(t *testing.T) {
	t.Run("plain message", func(t *testing.T) {
		err := ThrowException("ERR-001", "error message")
		exception := FromError(err)

		assert.Equal(t, "ERR-001", exception.Code)
		assert.Equal(t, err.Error(), "error message")
	})
	t.Run("message with parameters", func(t *testing.T) {
		err := ThrowException("ERR-002", "missing field %s", "field1")
		exception := FromError(err)

		assert.Equal(t, "ERR-002", exception.Code)
		assert.Equal(t, err.Error(), fmt.Sprintf("missing field %s", "field1"))
	})
}
