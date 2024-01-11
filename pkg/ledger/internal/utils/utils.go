package utils

import (
	"fmt"
	"os/exec"
)

// ExecuteShellCommand executes a shell command and returns the output and any error encountered.
func ExecuteShellCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute '%s' command: %w", command, err)
	}

	return string(output), nil
}
