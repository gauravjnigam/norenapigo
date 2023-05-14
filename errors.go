package norenapigo

// Error is the error type used for all API errors.
type Error struct {
	Code    string
	Message string
	Data    interface{}
}

// This makes Error a valid Go error type.
func (e Error) Error() string {
	return e.Message
}

// NewError creates and returns a new instace of Error
// with custom error metadata.
func NewError(etype string, message string, data interface{}) error {
	err := Error{}
	err.Message = message
	err.Code = etype
	err.Data = data
	return err
}
