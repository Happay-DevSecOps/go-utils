# go-utils
Common Golang utilities

### Systematic Error/Output handling (cerrors/couts) - 

##### define structs
```
// GeneralError struct (500-510)
type GeneralError struct {
	Code    string
	Message string
	Err     error
}

func (c *GeneralError) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s, Err: %s", c.Code, c.Message, c.Err)
}

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
```

##### Use structs
```
// Sample snippet to use Custom Error
func openFile() (Output string, out interface{}) {
	var file *os.File
	if file, err = os.Open("test.txt"); err != nil {
		generalError := &GeneralError{Code: "500", Message: "Unable to open file", Err: err}
		return "Failed", generalError
	}

	generalOutput := &GeneralOutput{Code: "200", Message: "File opened successfully", Out: file.Name()}

	return "Sucess", generalOutput
}

func main() {
	output, out := openFile()
	// Check error type
	if out, ok := out.(*cerrors.GeneralError); ok {
		finalError := output + ": GeneralError - " + out.Error()
		log.Println(finalError)
	}
	if out, ok := out.(*couts.GeneralOutput); ok {
		finalOutput := output + ": GeneralOutput - " + fmt.Sprintf("%v", out)
		log.Println(finalOutput)
	}
}
```
Output - 
```
2021/03/08 13:47:59 Failed: GeneralError - Code: 500, Message: Unable to open file, Err: open test.txt: no such file or directory
2021/03/08 13:47:60 Success: GeneralOutput - Code: 200, Message: file opened successfully, Out: test.txt
```
