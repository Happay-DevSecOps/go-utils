package couts

import "fmt"

// GeneralOutput struct (5)
type GeneralOutput struct {
	Code    string
	Message string
	Out     string
}

//Outp method
func (c *GeneralOutput) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s, Out: %s", c.Code, c.Message, c.Out)
}
