package login

import "fmt"

type ConfigFileLoadError struct {
	Err error
}

func (e ConfigFileLoadError) Error() string {
	return fmt.Errorf("failed to load default config: %w", e.Err).Error()
}

func NewConfigFileLoadError(err error) ConfigFileLoadError {
	return ConfigFileLoadError{
		Err: err,
	}
}
