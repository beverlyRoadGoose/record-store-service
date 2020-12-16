package errors

import (
	"fmt"
)

type BadRequest struct {
	Message string
}

func (e *BadRequest) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

type PreconditionFailed struct {
	Message string
}

func (e *PreconditionFailed) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
