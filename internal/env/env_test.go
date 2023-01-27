//go:build !integration

package env

import (
	"os"
	"testing"

	"github.com/jannotti-glaucio/timescale-assignment/internal/excepts"
	"github.com/stretchr/testify/assert"
)

const testFilePath = "./input.csv"
const testDbUrl = "postgres://user:passwd@host:5432/homework"

func TestCheckVars(t *testing.T) {
	t.Run("sucessfull", func(t *testing.T) {
		os.Setenv(FilePath, testFilePath)
		os.Setenv(DbUrl, testDbUrl)

		err := CheckVars()

		assert.Nil(t, err)
	})
	t.Run("missing "+FilePath+" variable", func(t *testing.T) {
		os.Setenv(FilePath, "")
		os.Setenv(DbUrl, testDbUrl)

		err := CheckVars()
		exception := excepts.FromError(err)

		assert.NotNil(t, err)
		assert.Equal(t, exception.Code, excepts.MissingEnvVariable)
	})
	t.Run("missing "+DbUrl+" variable", func(t *testing.T) {
		os.Setenv(FilePath, testFilePath)
		os.Setenv(DbUrl, "")

		err := CheckVars()
		exception := excepts.FromError(err)

		assert.NotNil(t, err)
		assert.Equal(t, exception.Code, excepts.MissingEnvVariable)
	})
}
