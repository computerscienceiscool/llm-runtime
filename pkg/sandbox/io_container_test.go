package sandbox

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestRunIOContainer_SimpleCommand(t *testing.T) {
	if !isDockerAvailable() {
		t.Skip("Docker not available")
	}

	tempDir := t.TempDir()
	output, err := RunIOContainer(
		tempDir,
		"alpine:latest",
		"echo hello",
		5*time.Second,
		"128m",
		1,
	)
	if err != nil {
		t.Fatalf("RunIOContainer() error = %v", err)
	}
	if !strings.Contains(output, "hello") {
		t.Errorf("Expected 'hello' in output, got: %s", output)
	}
}

func TestRunIOContainer_Timeout(t *testing.T) {
	if !isDockerAvailable() {
		t.Skip("Docker not available")
	}

	tempDir := t.TempDir()
	_, err := RunIOContainer(
		tempDir,
		"alpine:latest",
		"sleep 10",
		100*time.Millisecond,
		"128m",
		1,
	)
	if err == nil {
		t.Error("Expected timeout error")
	}
	// Accept either timeout message or context deadline exceeded
	if !strings.Contains(err.Error(), "timed out") && !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Errorf("Expected timeout error, got: %v", err)
	}
}

func TestReadFileInContainer_Success(t *testing.T) {
	if !isDockerAvailable() {
		t.Skip("Docker not available")
	}

	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.txt")
	testContent := "file content here"
	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	content, err := ReadFileInContainer(
		testFile,
		tempDir,
		"alpine:latest",
		5*time.Second,
		"128m",
		1,
	)
	if err != nil {
		t.Fatalf("ReadFileInContainer() error = %v", err)
	}
	if !strings.Contains(content, testContent) {
		t.Errorf("Expected %q in output, got: %s", testContent, content)
	}
}

func TestWriteFileInContainer_Success(t *testing.T) {
	if !isDockerAvailable() {
		t.Skip("Docker not available")
	}

	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "output.txt")
	testContent := "written by container"

	err := WriteFileInContainer(
		testFile,
		testContent,
		tempDir,
		"alpine:latest",
		5*time.Second,
		"128m",
		1,
	)
	if err != nil {
		t.Fatalf("WriteFileInContainer() error = %v", err)
	}

	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read written file: %v", err)
	}
	if string(content) != testContent {
		t.Errorf("Expected %q, got: %q", testContent, string(content))
	}
}

func TestEnsureIOContainerImage_AlpineExists(t *testing.T) {
	if !isDockerAvailable() {
		t.Skip("Docker not available")
	}

	err := EnsureIOContainerImage("alpine:latest")
	if err != nil {
		t.Errorf("EnsureIOContainerImage(alpine:latest) error = %v", err)
	}
}

func TestValidateIOContainer_Success(t *testing.T) {
	if !isDockerAvailable() {
		t.Skip("Docker not available")
	}

	tempDir := t.TempDir()
	err := ValidateIOContainer(tempDir, "alpine:latest")
	if err != nil {
		t.Errorf("ValidateIOContainer() error = %v", err)
	}
}

func TestValidateIOContainer_NonExistentRepo(t *testing.T) {
	if !isDockerAvailable() {
		t.Skip("Docker not available")
	}

	err := ValidateIOContainer("/nonexistent/path/12345", "alpine:latest")
	if err == nil {
		t.Error("Expected error for non-existent repository")
	}
}

// TestParseMemoryLimitIO tests memory limit string parsing for IO containers
func TestParseMemoryLimitIO(t *testing.T) {
	tests := []struct {
		name      string
		limit     string
		expected  int64
		expectErr bool
	}{
		{"empty string", "", 0, false},
		{"128 megabytes lowercase", "128m", 128 * 1024 * 1024, false},
		{"128 megabytes uppercase", "128M", 128 * 1024 * 1024, false},
		{"512 megabytes", "512m", 512 * 1024 * 1024, false},
		{"1 gigabyte lowercase", "1g", 1 * 1024 * 1024 * 1024, false},
		{"1 gigabyte uppercase", "1G", 1 * 1024 * 1024 * 1024, false},
		{"2 gigabytes", "2g", 2 * 1024 * 1024 * 1024, false},
		{"invalid format", "invalid", 0, true},
		{"no suffix", "256", 0, true},
		{"zero megabytes", "0m", 0, true},
		{"negative megabytes", "-128m", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseMemoryLimit(tt.limit)
			if tt.expectErr {
				if err == nil {
					t.Errorf("parseMemoryLimit(%q) expected error, got nil", tt.limit)
				}
			} else {
				if err != nil {
					t.Errorf("parseMemoryLimit(%q) unexpected error: %v", tt.limit, err)
				}
				if result != tt.expected {
					t.Errorf("parseMemoryLimit(%q) = %d, want %d", tt.limit, result, tt.expected)
				}
			}
		})
	}
}

// BenchmarkReadFile_Native benchmarks direct file reading
func BenchmarkReadFile_Native(b *testing.B) {
	tmpDir := b.TempDir()

	// Create test file
	testFile := filepath.Join(tmpDir, "bench.txt")
	testContent := strings.Repeat("benchmark content\n", 100)
	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		b.Fatalf("failed to create test file: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := os.ReadFile(testFile)
		if err != nil {
			b.Fatalf("read failed: %v", err)
		}
	}
}

// BenchmarkReadFile_Containerized benchmarks containerized file reading
func BenchmarkReadFile_Containerized(b *testing.B) {
	if !isDockerAvailable() {
		b.Skip("Docker not available")
	}

	tmpDir := b.TempDir()

	// Create test file
	testFile := filepath.Join(tmpDir, "bench.txt")
	testContent := strings.Repeat("benchmark content\n", 100)
	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		b.Fatalf("failed to create test file: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ReadFileInContainer(
			testFile,
			tmpDir,
			"alpine:latest",
			5*time.Second,
			"128m",
			1,
		)
		if err != nil {
			b.Fatalf("containerized read failed: %v", err)
		}
	}
}

// BenchmarkWriteFile_Native benchmarks direct file writing
func BenchmarkWriteFile_Native(b *testing.B) {
	tmpDir := b.TempDir()
	testContent := strings.Repeat("benchmark content\n", 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testFile := filepath.Join(tmpDir, fmt.Sprintf("bench_%d.txt", i))
		err := os.WriteFile(testFile, []byte(testContent), 0644)
		if err != nil {
			b.Fatalf("write failed: %v", err)
		}
	}
}

// BenchmarkWriteFile_Containerized benchmarks containerized file writing
func BenchmarkWriteFile_Containerized(b *testing.B) {
	if !isDockerAvailable() {
		b.Skip("Docker not available")
	}

	tmpDir := b.TempDir()
	testContent := strings.Repeat("benchmark content\n", 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testFile := filepath.Join(tmpDir, fmt.Sprintf("bench_%d.txt", i))
		err := WriteFileInContainer(
			testFile,
			testContent,
			tmpDir,
			"alpine:latest",
			5*time.Second,
			"128m",
			1,
		)
		if err != nil {
			b.Fatalf("containerized write failed: %v", err)
		}
	}
}

// BenchmarkReadFile_SmallFile_Native benchmarks native read of small file
func BenchmarkReadFile_SmallFile_Native(b *testing.B) {
	tmpDir := b.TempDir()
	testFile := filepath.Join(tmpDir, "small.txt")
	if err := os.WriteFile(testFile, []byte("small"), 0644); err != nil {
		b.Fatalf("failed to create test file: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		os.ReadFile(testFile)
	}
}

// BenchmarkReadFile_SmallFile_Containerized benchmarks containerized read of small file
func BenchmarkReadFile_SmallFile_Containerized(b *testing.B) {
	if !isDockerAvailable() {
		b.Skip("Docker not available")
	}

	tmpDir := b.TempDir()
	testFile := filepath.Join(tmpDir, "small.txt")
	if err := os.WriteFile(testFile, []byte("small"), 0644); err != nil {
		b.Fatalf("failed to create test file: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReadFileInContainer(testFile, tmpDir, "alpine:latest", 5*time.Second, "128m", 1)
	}
}

// BenchmarkReadFile_LargeFile_Native benchmarks native read of large file
func BenchmarkReadFile_LargeFile_Native(b *testing.B) {
	tmpDir := b.TempDir()
	testFile := filepath.Join(tmpDir, "large.txt")
	largeContent := strings.Repeat("x", 100*1024) // 100KB
	if err := os.WriteFile(testFile, []byte(largeContent), 0644); err != nil {
		b.Fatalf("failed to create test file: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		os.ReadFile(testFile)
	}
}

// BenchmarkReadFile_LargeFile_Containerized benchmarks containerized read of large file
func BenchmarkReadFile_LargeFile_Containerized(b *testing.B) {
	if !isDockerAvailable() {
		b.Skip("Docker not available")
	}

	tmpDir := b.TempDir()
	testFile := filepath.Join(tmpDir, "large.txt")
	largeContent := strings.Repeat("x", 100*1024) // 100KB
	if err := os.WriteFile(testFile, []byte(largeContent), 0644); err != nil {
		b.Fatalf("failed to create test file: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReadFileInContainer(testFile, tmpDir, "alpine:latest", 5*time.Second, "128m", 1)
	}
}

func TestReadDockerLogs_NilStdout(t *testing.T) {
	err := readDockerLogs(strings.NewReader(""), nil)
	if err == nil {
		t.Error("expected error for nil stdout")
	}
	if !strings.Contains(err.Error(), "must not be nil") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestReadDockerLogs_EmptyStream(t *testing.T) {
	var stdout strings.Builder
	err := readDockerLogs(strings.NewReader(""), &stdout)
	if err != nil {
		t.Errorf("empty stream should return nil, got: %v", err)
	}
}

func TestReadDockerLogs_WriteError(t *testing.T) {
	// Build a valid Docker log frame: header(8 bytes) + payload
	header := []byte{1, 0, 0, 0, 0, 0, 0, 5}
	payload := []byte("hello")
	frame := append(header, payload...)

	stdout := &ioErrWriter{err: fmt.Errorf("disk full")}
	err := readDockerLogs(strings.NewReader(string(frame)), stdout)
	if err == nil {
		t.Error("expected error when stdout write fails")
	}
	if !strings.Contains(err.Error(), "disk full") {
		t.Errorf("expected 'disk full' in error, got: %v", err)
	}
}

// ioErrWriter is an io.Writer that always returns an error.
type ioErrWriter struct {
	err error
}

func (w *ioErrWriter) Write(p []byte) (int, error) {
	return 0, w.err
}
