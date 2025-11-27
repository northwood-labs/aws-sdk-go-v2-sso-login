package login

import "fmt"

type SsoOidcTokenCreationError struct {
	Err error
}

func (e SsoOidcTokenCreationError) Error() string {
	return fmt.Errorf("failed to retrieve user: %w", e.Err).Error()
}

func NewSsoOidcTokenCreationError(err error) SsoOidcTokenCreationError {
	return SsoOidcTokenCreationError{
		Err: err,
	}
}
