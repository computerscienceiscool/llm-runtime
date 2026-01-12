package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type projectHint struct {
	marker string
	config string
	note   string
}

var hints = []projectHint{
	{marker: "go.mod", config: "go-developer.yaml", note: "Go project detected"},
	{marker: "package.json", config: "node-developer.yaml", note: "Node.js project detected"},
	{marker: "requirements.txt", config: "python-developer.yaml", note: "Python project detected"},
	{marker: "pyproject.toml", config: "python-developer.yaml", note: "Python project detected"},
	{marker: "Gemfile", config: "ruby-developer.yaml", note: "Ruby project detected"},
}

func runAutoDetect() error {
	root, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("auto-detect: cannot determine working directory: %w", err)
	}

	var found []string
	for _, h := range hints {
		if fileExists(filepath.Join(root, h.marker)) {
			msg := h.note
			if h.config != "" {
				msg += fmt.Sprintf(" → Suggested config: %s", h.config)
			}
			found = append(found, msg)
		}
	}

	if len(found) == 0 {
		fmt.Println("auto-detect: no known project markers found")
		fmt.Println("You can specify --root and a config manually.")
		return nil
	}

	fmt.Println("auto-detect results:")
	for _, msg := range found {
		fmt.Println(" - " + msg)
	}

	return nil
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
