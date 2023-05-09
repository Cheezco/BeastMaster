package internal

import (
	"bufio"
	"os/exec"
)

func RunExecutableWithScanner(path string, args ...string) (*bufio.Scanner, *exec.Cmd) {
	cmd := exec.Command(path, args...)

	stdout, err := cmd.StdoutPipe()
	CheckError(err)
	err = cmd.Start()
	CheckError(err)

	return bufio.NewScanner(stdout), cmd
}
