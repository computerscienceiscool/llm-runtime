package sandbox

import (
	"fmt"
	"github.com/computerscienceiscool/llm-runtime/pkg/config"
	"strings"
)

// ValidateExecCommand checks if the command is allowed to execute
// Note: Exec is always enabled in container-only mode
func ValidateExecCommand(command string, whitelist []string) error {
	// Trim whitespace and validate input
	command = strings.TrimSpace(command)

	// Check for empty command
	if command == "" {
		return fmt.Errorf("empty command")
	}

	// Check command length (prevent abuse with extremely long commands)
	const maxCommandLength = config.MaxCommandLength
	if len(command) > maxCommandLength {
		return fmt.Errorf("command too long (max %d characters, got %d)", maxCommandLength, len(command))
	}

	// Check for null bytes or other control characters
	if strings.ContainsAny(command, "\x00\x01\x02\x03\x04\x05\x06\x07\x08") {
		return fmt.Errorf("command contains invalid control characters")
	}

	// Check whitelist is not empty
	if len(whitelist) == 0 {
		return fmt.Errorf("no commands are whitelisted")
	}

	// Normalize command and whitelist by collapsing whitespace
	command = strings.Join(strings.Fields(command), " ")

	// Check against whitelist using exact match or full-token prefix (next char must be space or end)
	for _, allowed := range whitelist {
		allowedNorm := strings.Join(strings.Fields(allowed), " ")
		if allowedNorm == "" {
			continue
		}
		if command == allowedNorm {
			return nil
		}
		if strings.HasPrefix(command, allowedNorm+" ") {
			return nil
		}
	}

	return fmt.Errorf("command not in whitelist: %s", command)
}
