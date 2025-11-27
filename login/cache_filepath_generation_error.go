package login

import "fmt"

type CacheFilepathGenerationError struct {
	Err      error
	Profile  string
	SSOStart string
}

func (e CacheFilepathGenerationError) Error() string {
	return fmt.Errorf(
		"failed to generate cache file for profile %q with URL %s",
		e.Profile,
		e.SSOStart,
	).Error()
}

func NewCacheFilepathGenerationError(err error, profile, ssoStart string) CacheFilepathGenerationError {
	return CacheFilepathGenerationError{
		Err:      err,
		Profile:  profile,
		SSOStart: ssoStart,
	}
}
