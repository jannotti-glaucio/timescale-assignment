//go:build !integration

package file

import (
	"testing"

	"github.com/jannotti-glaucio/timescale-assignment/internal/tests"
	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {

	t.Run("exists", func(t *testing.T) {
		exists := FileExists(tests.TestDataPath + "env/file_exists.txt")

		assert.True(t, exists)
	})
	t.Run("not exists", func(t *testing.T) {
		exists := FileExists("fileNotFound.txt")

		assert.False(t, exists)
	})
}
