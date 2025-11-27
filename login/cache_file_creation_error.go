package login

import "fmt"

type CacheFileCreationError struct {
	Err  error
	Path string
}

func (e CacheFileCreationError) Error() string {
	return fmt.Errorf("cache file %s creation failed: %w", e.Path, e.Err).Error()
}

func NewCacheFileCreationError(err error, cacheFilePath string) CacheFileCreationError {
	return CacheFileCreationError{
		Err:  err,
		Path: cacheFilePath,
	}
}
