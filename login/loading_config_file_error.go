package login

import "fmt"

type LoadingConfigFileError struct {
	Err            error
	ConfigFilePath string
}

func (e LoadingConfigFileError) Error() string {
	return fmt.Errorf("failed to load config file %q: %w", e.ConfigFilePath, e.Err).Error()
}

func NewLoadingConfigFileError(err error, configFilePath string) LoadingConfigFileError {
	return LoadingConfigFileError{
		ConfigFilePath: configFilePath,
		Err:            err,
	}
}
