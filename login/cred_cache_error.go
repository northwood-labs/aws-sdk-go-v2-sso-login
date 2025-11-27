package login

import "fmt"

type CredCacheError struct {
	Err error
}

func (e CredCacheError) Error() string {
	return fmt.Errorf("failed to retrieve creds").Error()
}

func NewCredCacheError(err error) CredCacheError {
	return CredCacheError{
		Err: err,
	}
}
