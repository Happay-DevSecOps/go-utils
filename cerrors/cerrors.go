package cerrors

import "fmt"

// GeneralError struct
type GeneralError struct {
	Code    string
	Message string
	Err     error
}

func (c *GeneralError) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s, Error: %s", c.Code, c.Message, c.Err)
}
