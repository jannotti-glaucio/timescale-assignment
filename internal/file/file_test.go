package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {

	t.Run("exists", func(t *testing.T) {
		exists := FileExists("../../test/data/env/file.txt")

		assert.True(t, exists)
	})
	t.Run("not exists", func(t *testing.T) {
		exists := FileExists("fileNotFound.txt")

		assert.False(t, exists)
	})
}
