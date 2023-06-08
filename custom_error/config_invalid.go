package custom_error

import "fmt"

type InvalidErr struct {
	Name  string
	Field string
	Value interface{}
}

func (n InvalidErr) Error() string {
	if n.Name != "" {
		return fmt.Sprintf("%s.%s(%v) is invalid", n.Name, n.Field, n.Value)
	}
	return fmt.Sprintf("%s(%v) is invalid", n.Field, n.Value)
}
