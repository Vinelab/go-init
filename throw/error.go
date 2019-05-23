package throw

type BaseError struct {
	Message string
	Code    uint16
}

func (e BaseError) Error() string {
	return e.Message
}
