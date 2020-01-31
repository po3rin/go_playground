package errs

type ErrorUnknown struct {
}

func (e *ErrorUnknown) Error() string {
	return ""
}
func (e *ErrorUnknown) StatusCode() int {
	return 0
}

type ErrorResourceNotFound struct {
}

func (e *ErrorResourceNotFound) Error() string {
	return ""
}
func (e *ErrorResourceNotFound) StatusCode() int {
	return 0
}
