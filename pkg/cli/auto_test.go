package cli

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunAutoDetect_NoMarkers(t *testing.T) {
	dir := t.TempDir()
	orig, _ := os.Getwd()
	defer os.Chdir(orig)
	os.Chdir(dir)

	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	if err := runAutoDetect(); err != nil {
		t.Fatalf("runAutoDetect() error = %v", err)
	}

	w.Close()
	os.Stdout = stdout
	out, _ := io.ReadAll(r)
	buf.Write(out)

	output := buf.String()
	if !strings.Contains(output, "no known project markers") {
		t.Errorf("expected no markers message, got: %s", output)
	}
}

func TestRunAutoDetect_WithMarkers(t *testing.T) {
	dir := t.TempDir()
	orig, _ := os.Getwd()
	defer os.Chdir(orig)
	os.Chdir(dir)

	markers := []string{"go.mod", "package.json", "pyproject.toml"}
	for _, m := range markers {
		if err := os.WriteFile(filepath.Join(dir, m), []byte(""), 0644); err != nil {
			t.Fatalf("failed to create marker %s: %v", m, err)
		}
	}

	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	if err := runAutoDetect(); err != nil {
		t.Fatalf("runAutoDetect() error = %v", err)
	}

	w.Close()
	os.Stdout = stdout
	out, _ := io.ReadAll(r)
	buf.Write(out)

	output := buf.String()
	for _, expect := range []string{"Go project detected", "Node.js project detected", "Python project detected"} {
		if !strings.Contains(output, expect) {
			t.Errorf("expected output to contain %q, got: %s", expect, output)
		}
	}
	if !strings.Contains(output, "preview") {
		t.Errorf("expected output to mention preview, got: %s", output)
	}
}
