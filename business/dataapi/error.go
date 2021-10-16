package dataapi

// Error implements the error interface and is used to identify a trusted error
type Error struct {
	Err error
}

func (err *Error) Error() string {
	return err.Err.Error()
}
