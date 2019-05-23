package throw

// InvalidInput specifies that the received input parameters
type InvalidInput struct {
	*BaseError
}

func (e InvalidInput) Error() string {
	return "Invalid input error: " + e.BaseError.Error()
}

func InvalidInputError(m string, c uint16) InvalidInput {
	return InvalidInput{&BaseError{Message: m, Code: c}}
}
