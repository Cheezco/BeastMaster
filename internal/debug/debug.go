package debug

import (
	"os"
	"path/filepath"
)

func IsDebugMode() bool {
	return os.Getenv("DEBUG") == "true"
}

func GetExecutableName() string {
	executablePath, err := os.Executable()
	if err != nil {
		return "None"
	}

	executableName := filepath.Base(executablePath)
	return executableName[:len(executableName)-len(filepath.Ext(executableName))]
}
