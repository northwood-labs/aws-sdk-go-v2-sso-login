package login

// ProfileValidationError error validating the given AWS profile. A required
// value may be missing.
type ProfileValidationError struct {
	Profile string
	Path    string
	Field   string
	CurVal  string
	ExpVal  string
}

func (e ProfileValidationError) Error() string {
	return "Profile validation failed. " +
		"; Profile: " + e.Profile +
		"; Config file: " + e.Path +
		"; Field: " + e.Field +
		"; Value: " + e.CurVal +
		"; Expected " + e.ExpVal
}

func NewProfileValidationError(profile, path, field, curVal, expVal string) ProfileValidationError {
	return ProfileValidationError{profile, path, field, curVal, expVal}
}
