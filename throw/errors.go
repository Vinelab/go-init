package throw

// Error is what we should be handling throught the application
type Error interface {
	error
}

// HTTPError is the base error for the errors that are
// thrown among the HTTP context
type HTTPError struct {
	Message string
	Code    uint16
	Status  uint16
}

// InvalidInput specifies that the received input parameters
// are not as expected and have failed the validation
type InvalidInput struct {
	*HTTPError
}

// Error is implemented to comply with the "error" interface
// specified by Exception.
func (e InvalidInput) Error() string {
	return e.Message
}

// InvalidInputError creates and returns a new InvalidInput instance
// with the given data.
func InvalidInputError(m string, c, s uint16) InvalidInput {
	return InvalidInput{&HTTPError{Message: m, Code: c, Status: s}}
}
