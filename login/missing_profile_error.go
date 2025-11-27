package login

import "fmt"

type MissingProfileError struct {
	Err     error
	Profile string
	Path    string
}

func (e MissingProfileError) Error() string {
	return fmt.Errorf(
		"profile %s does not exist in config file %s",
		e.Profile,
		e.Path,
	).Error()
}

func NewMissingProfileError(err error, profile string, path string) MissingProfileError {
	return MissingProfileError{
		Err:     err,
		Profile: profile,
		Path:    path,
	}
}
