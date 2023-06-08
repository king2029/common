package custom_error

import "fmt"

type NotSetErr struct {
	Name  string
	Field string
}

func (n NotSetErr) Error() string {
	if n.Name != "" {
		return fmt.Sprintf("%s.%s not set", n.Name, n.Field)
	}
	return fmt.Sprintf("%s not set", n.Field)
}
