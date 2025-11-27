package login

import "fmt"

type OsUserError struct {
	Err error
}

func (e OsUserError) Error() string {
	return fmt.Errorf("failed to retrieve user: %w", e.Err).Error()
}

func NewOsUserError(err error) OsUserError {
	return OsUserError{
		Err: err,
	}
}
