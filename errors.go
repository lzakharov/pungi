package pungi

import "fmt"

const errInvalidFormat = "'%s' has zero value"

// ErrInvalid indicates zero value in object.
type ErrInvalid struct {
	Name string
}

func (e *ErrInvalid) Error() string {
	return fmt.Sprintf(errInvalidFormat, e.Name)
}
