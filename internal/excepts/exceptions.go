package excepts

import (
	"fmt"
)

type Exception struct {
	Code    string
	Message string
}

func (e Exception) Error() string {
	return e.Message
}

func FromError(err error) Exception {
	return err.(Exception)
}

func ThrowException(code string, message string, args ...interface{}) error {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	return Exception{
		Code:    code,
		Message: message,
	}
}

const (
	ErrorConfiguringLogger  = "ERR-LOG-001"
	ErrorConnectingToDB     = "ERR-DB-001"
	ErrorExecutingQuery     = "ERR-DB-002"
	ErrorDisconectingFromDB = "ERR-DB-003"
	ErrorLoadingEnvFile     = "ERR-ENV-001"
	MissingEnvVariable      = "ERR-ENV-002"
	FailedOpeningCSVFile    = "ERR-PARSE-001"
	InvalidNumberOfColumns  = "ERR-PARSE-002"
	MissingHostnameColumn   = "ERR-PARSE-003"
	MissingStartDateColumn  = "ERR-PARSE-004"
	InvalidStartDateColumn  = "ERR-PARSE-005"
	MissingEndDateColumn    = "ERR-PARSE-006"
	InvalidEndDateColumn    = "ERR-PARSE-007"
)
