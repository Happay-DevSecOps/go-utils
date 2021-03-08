package cerrors

import "fmt"

// GeneralError struct (500-510)
type GeneralError struct {
	Code    string
	Message string
	Err     error
}

func (c *GeneralError) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s, Err: %s", c.Code, c.Message, c.Err)
}
