package pungi

import "fmt"

const errInvalidFormat = "'%s' has a zero value"

// ErrInvalid indicates a zero value in an object.
type ErrInvalid struct {
	Name string
}

func (e *ErrInvalid) Error() string {
	return fmt.Sprintf(errInvalidFormat, e.Name)
}
