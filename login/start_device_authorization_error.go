package login

import "fmt"

type StartDeviceAuthorizationError struct {
	Err error
}

func (e StartDeviceAuthorizationError) Error() string {
	return fmt.Errorf("failed to start device authorization: %w", e.Err).Error()
}

func NewStartDeviceAuthorizationError(err error) StartDeviceAuthorizationError {
	return StartDeviceAuthorizationError{
		Err: err,
	}
}
