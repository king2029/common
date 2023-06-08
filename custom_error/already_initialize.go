package custom_error

import "fmt"

type AlreadyInitializeErr struct {
	Name string
}

func (n AlreadyInitializeErr) Error() string {
	return fmt.Sprintf("%s already initialize", n.Name)
}
