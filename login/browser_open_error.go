package login

import "fmt"

type BrowserOpenError struct {
	Err error
}

func (e BrowserOpenError) Error() string {
	return fmt.Errorf("failed to open a browser: %w", e.Err).Error()
}

func NewBrowserOpenError(err error) BrowserOpenError {
	return BrowserOpenError{
		Err: err,
	}
}
