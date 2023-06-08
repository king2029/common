package custom_error

import "fmt"

type RegisterModelErr struct {
	Package string
	Name    string
	Err     error
}

func (e RegisterModelErr) Error() string {
	return fmt.Sprintf("%s register model(%s) error: %s", e.Package, e.Name, e.Err)
}
