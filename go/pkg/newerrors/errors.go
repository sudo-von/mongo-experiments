package newerrors

// ErrNotFound Error when not found something in the Database.
type ErrNotFound struct {
	message string
}

func (e *ErrNotFound) Error() string {
	return e.message
}

// NewErrNotFound creates a ErrNotFound error.
func NewErrNotFound(message string) error {
	return &ErrNotFound{
		message: message,
	}
}

// ErrUnauthorized Error when not found something in the Database.
type ErrUnauthorized struct {
	message string
}

func (e *ErrUnauthorized) Error() string {
	return e.message
}

// NewErrUnauthorized creates a ErrUnauthorized error.
func NewErrUnauthorized(message string) error {
	return &ErrUnauthorized{
		message: message,
	}
}

// ErrBadRequest Error when is bad request.
type ErrBadRequest struct {
	message string
}

func (e *ErrBadRequest) Error() string {
	return e.message
}

// NewErrBadRequest creates a ErrBadRequest error.
func NewErrBadRequest(message string) error {
	return &ErrBadRequest{
		message: message,
	}
}

// ErrConflict Error when not found something in the Database.
type ErrConflict struct {
	message string
}

func (e *ErrConflict) Error() string {
	return e.message
}

// NewErrConflict creates a ErrConflict error.
func NewErrConflict(message string) error {
	return &ErrConflict{
		message: message,
	}
}

// ErrForbidden Error indicates that the server understood the request but refuses to authorize it.
type ErrForbidden struct {
	message string
}

func (e *ErrForbidden) Error() string {
	return e.message
}

// NewErrForbidden creates a ErrForbidden error.
func NewErrForbidden(message string) error {
	return &ErrForbidden{
		message: message,
	}
}
