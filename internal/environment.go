package internal

import (
	"os"
	"runtime"
)

func IsInsideDockerContainer() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	return false
}

//goland:noinspection GoBoolExpressions
func IsWindows() bool {
	return runtime.GOOS == "windows"
}
