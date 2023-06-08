package custom_error

import "fmt"

type RecordExistErr struct {
	Table string
	Field string
	Value interface{}
}

func (n RecordExistErr) Error() string {
	return fmt.Sprintf("%s.%s(%v) already exist", n.Table, n.Field, n.Value)
}
