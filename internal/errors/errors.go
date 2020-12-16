package errors

import (
	"fmt"
)

type BadRequest struct {
	Message string
}

func (r *BadRequest) Error() string {
	return fmt.Sprintf("%s", r.Message)
}

type PreconditionFailed struct {
	Message string
}

func (r *PreconditionFailed) Error() string {
	return fmt.Sprintf("%s", r.Message)
}
