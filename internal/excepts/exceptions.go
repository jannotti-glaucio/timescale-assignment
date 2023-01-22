package excepts

import (
	"fmt"
)

type Exception struct {
	Code    string
	Message string
}

func (e *Exception) Error() string {
	return e.Message
}

func ThrowException(code string, message string, args ...interface{}) *Exception {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	return &Exception{
		Code:    code,
		Message: message,
	}
}

func RethrowException(err *Exception, message string, args ...interface{}) *Exception {
	return ThrowException(err.Code, message, args...)
}

const (
	ErrorConnectingToDB    = "ERR-DB-001"
	ErrorExecutingQuery    = "ERR-DB-002"
	ErrorLoadingEnvFile    = "ERR-ENV-001"
	MissingEnvVariable     = "ERR-ENV-002"
	FailedOpeningCSVFile   = "ERR-PARSE-001"
	InvalidNumberOfColumns = "ERR-PARSE-002"
	MissingHostnameColumn  = "ERR-PARSE-003"
	MissingStartDateColumn = "ERR-PARSE-004"
	InvalidStartDateColumn = "ERR-PARSE-005"
	MissingEndDateColumn   = "ERR-PARSE-006"
	InvalidEndDateColumn   = "ERR-PARSE-007"
)
