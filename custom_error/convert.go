package custom_error

import "fmt"

type ConvertErr struct {
	Name string
	Type string
	Err  error
}

func (c ConvertErr) Error() string {
	return fmt.Sprintf("%s convert to %s error: %s", c.Name, c.Type, c.Err)
}
