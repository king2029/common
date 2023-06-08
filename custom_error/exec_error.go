package custom_error

import "fmt"

type ExecErr struct {
	Output []byte
	Err    error
}

func (e ExecErr) Error() string {
	return fmt.Sprintf("exec error: %s | %s", e.Output, e.Err)
}
