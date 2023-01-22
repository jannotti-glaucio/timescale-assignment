package env

import (
	"fmt"
	"os"
	"testing"

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

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), fmt.Sprintf(errorMissingVariable, FilePath))
	})
	t.Run("missing "+DbUrl+" variable", func(t *testing.T) {
		os.Setenv(FilePath, testFilePath)
		os.Setenv(DbUrl, "")

		err := CheckVars()

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), fmt.Sprintf(errorMissingVariable, DbUrl))
	})
}
