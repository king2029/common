package common

import (
	"errors"
	"os/exec"
)

func BashExecute(command string) (string, error) {
	output, err := exec.Command("/bin/bash", "-c", command).CombinedOutput()
	if err != nil {
		return "", errors.New(string(output))
	}

	return string(output), nil
}
