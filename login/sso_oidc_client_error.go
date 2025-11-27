package login

import "fmt"

type SsoOidcClientError struct {
	Err error
}

func (e SsoOidcClientError) Error() string {
	return fmt.Errorf("failed to register ssoOidcClient: %w", e.Err).Error()
}

func NewSsoOidcClientError(err error) SsoOidcClientError {
	return SsoOidcClientError{
		Err: err,
	}
}
