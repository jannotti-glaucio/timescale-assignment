//go:build !integration

package logger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert.NotPanics(t, func() {
		Init()
	})
}

func TestInfo(t *testing.T) {
	assert.NotPanics(t, func() {
		Info("message")
	})
}

func TestDebug(t *testing.T) {
	assert.NotPanics(t, func() {
		Debug("message")
	})
}

func TestFormatMessage(t *testing.T) {
	t.Run("plain message", func(t *testing.T) {
		message := formatMessage("error message")

		assert.Equal(t, message, "error message")
	})
	t.Run("message with parameters", func(t *testing.T) {
		message := formatMessage("missing field %s", "field1")

		assert.Equal(t, message, fmt.Sprintf("missing field %s", "field1"))
	})
}
