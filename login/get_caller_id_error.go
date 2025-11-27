package login

type GetCallerIdError struct {
	Err error
}

func (e GetCallerIdError) Error() string {
	return "stsClient.GetCallerIdentity failed"
}

func NewGetCallerIdError(err error) GetCallerIdError {
	return GetCallerIdError{
		Err: err,
	}
}
