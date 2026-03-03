package sandbox

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
)

// RunIOContainer executes a containerized I/O operation
func RunIOContainer(repoRoot, containerImage, command string, timeout time.Duration, memLimit string, cpuLimit int) (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", fmt.Errorf("failed to create Docker client: %w", err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Configure container
	containerConfig := &container.Config{
		Image:      containerImage,
		Cmd:        strslice.StrSlice{"/bin/sh", "-c", command},
		WorkingDir: "/workspace",
		User:       "1000:1000",
	}

	// Parse memory limit
	memoryBytes, err := parseMemoryLimit(memLimit)
	if err != nil {
		return "", fmt.Errorf("invalid container config: %w", err)
	}

	// Configure host
	hostConfig := &container.HostConfig{
		NetworkMode: "none",
		Resources: container.Resources{
			Memory:   memoryBytes,
			NanoCPUs: int64(cpuLimit) * 1000000000,
		},
		Mounts: []mount.Mount{
			{
				Type:     mount.TypeBind,
				Source:   repoRoot,
				Target:   "/workspace",
				ReadOnly: true,
			},
		},
		CapDrop:     strslice.StrSlice{"ALL"},
		SecurityOpt: []string{"no-new-privileges"},
	}

	// Create container
	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		return "", fmt.Errorf("failed to create container: %w", err)
	}
	defer cli.ContainerRemove(ctx, resp.ID, container.RemoveOptions{Force: true})

	// Start container
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return "", fmt.Errorf("failed to start container: %w", err)
	}

	// Wait for completion
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return "", fmt.Errorf("container execution failed: %w", err)
		}
	case <-statusCh:
		// Container finished
	case <-ctx.Done():
		return "", fmt.Errorf("I/O operation timed out after %v", timeout)
	}

	// Get logs
	logReader, err := cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		return "", fmt.Errorf("failed to get container logs: %w", err)
	}
	defer logReader.Close()

	// Read output
	var stdout strings.Builder
	if err := readDockerLogs(logReader, &stdout); err != nil {
		return "", fmt.Errorf("failed to read container output: %w", err)
	}

	return stdout.String(), nil
}

// ReadFileInContainer reads a file using the I/O container
func ReadFileInContainer(filePath, repoRoot, containerImage string, timeout time.Duration, memLimit string, cpuLimit int) (string, error) {
	// Make path relative to repo root for container
	relPath, err := filepath.Rel(repoRoot, filePath)
	if err != nil {
		return "", fmt.Errorf("failed to get relative path: %w", err)
	}

	command := fmt.Sprintf("cat /workspace/%s", relPath)
	return RunIOContainer(repoRoot, containerImage, command, timeout, memLimit, cpuLimit)
}

// WriteFileInContainer writes a file using the I/O container
func WriteFileInContainer(filePath, content, repoRoot, containerImage string, timeout time.Duration, memLimit string, cpuLimit int) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %w", err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	relPath, err := filepath.Rel(repoRoot, filePath)
	if err != nil {
		return fmt.Errorf("failed to get relative path: %w", err)
	}

	// Write to temp file first, then move (atomic)
	// Directory creation happens inside container
	command := fmt.Sprintf("mkdir -p $(dirname /workspace/%s) && printf '%%s' %q > /workspace/%s.tmp && mv /workspace/%s.tmp /workspace/%s",
		relPath, content, relPath, relPath, relPath)

	// Configure container with read-write mount
	containerConfig := &container.Config{
		Image:      containerImage,
		Cmd:        strslice.StrSlice{"/bin/sh", "-c", command},
		WorkingDir: "/workspace",
		User:       "1000:1000",
	}

	memoryBytes, err := parseMemoryLimit(memLimit)
	if err != nil {
		return fmt.Errorf("invalid container config: %w", err)
	}

	hostConfig := &container.HostConfig{
		NetworkMode: "none",
		Resources: container.Resources{
			Memory:   memoryBytes,
			NanoCPUs: int64(cpuLimit) * 1000000000,
		},
		Mounts: []mount.Mount{
			{
				Type:     mount.TypeBind,
				Source:   repoRoot,
				Target:   "/workspace",
				ReadOnly: false, // Read-write for writes
			},
		},
		CapDrop:     strslice.StrSlice{"ALL"},
		SecurityOpt: []string{"no-new-privileges"},
	}

	// Create container
	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		return fmt.Errorf("failed to create container: %w", err)
	}
	defer cli.ContainerRemove(ctx, resp.ID, container.RemoveOptions{Force: true})

	// Start container
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return fmt.Errorf("failed to start container: %w", err)
	}

	// Wait for completion
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("container write failed: %w", err)
		}
	case <-statusCh:
		// Container finished
	case <-ctx.Done():
		return fmt.Errorf("write operation timed out after %v", timeout)
	}

	return nil
}

// EnsureIOContainerImage verifies the I/O container image exists
func EnsureIOContainerImage(imageName string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %w", err)
	}
	defer cli.Close()

	ctx := context.Background()
	_, _, err = cli.ImageInspectWithRaw(ctx, imageName)
	if err != nil {
		return fmt.Errorf("I/O container image not found: %s\nRun: docker build -f Dockerfile.io -t %s .", imageName, imageName)
	}
	return nil
}

// ValidateIOContainer runs pre-flight checks for containerized I/O
func ValidateIOContainer(repoRoot, containerImage string) error {
	// Check Docker is available
	if err := CheckDockerAvailability(); err != nil {
		return fmt.Errorf("Docker not available: %w", err)
	}

	// Check image exists
	if err := EnsureIOContainerImage(containerImage); err != nil {
		return err
	}

	// Check repo root exists and is readable
	if _, err := os.Stat(repoRoot); err != nil {
		return fmt.Errorf("repository root not accessible: %w", err)
	}

	return nil
}

// readDockerLogs reads Docker logs and extracts stdout.
// The stdout writer must be non-nil.
func readDockerLogs(reader io.Reader, stdout io.Writer) error {
	if stdout == nil {
		return fmt.Errorf("stdout writer must not be nil")
	}
	buf := make([]byte, 8)
	for {
		_, err := io.ReadFull(reader, buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		size := int(buf[4])<<24 | int(buf[5])<<16 | int(buf[6])<<8 | int(buf[7])
		if size > maxLogPayloadSize {
			return fmt.Errorf("Docker log frame too large: %d bytes (max %d)", size, maxLogPayloadSize)
		}
		payload := make([]byte, size)
		_, err = io.ReadFull(reader, payload)
		if err != nil {
			return err
		}

		if _, err := stdout.Write(payload); err != nil {
			return fmt.Errorf("failed to write output: %w", err)
		}
	}
}

// ReadFileInContainerPooled reads a file using a pooled container
func ReadFileInContainerPooled(ctx context.Context, pool *ContainerPool, filePath, repoRoot string) (string, error) {
	if pool == nil {
		// Fallback to non-pooled version
		return ReadFileInContainer(filePath, repoRoot, "llm-runtime-io:latest", 60*time.Second, "256m", 1)
	}

	relPath, err := filepath.Rel(repoRoot, filePath)
	if err != nil {
		return "", fmt.Errorf("failed to get relative path: %w", err)
	}

	command := fmt.Sprintf("cat /workspace/%s", relPath)
	return ExecuteInPooledContainer(ctx, pool, command, repoRoot)
}

// WriteFileInContainerPooled writes a file using a pooled container
func WriteFileInContainerPooled(ctx context.Context, pool *ContainerPool, filePath, content, repoRoot string) error {
	if pool == nil {
		// Fallback to non-pooled version
		return WriteFileInContainer(filePath, content, repoRoot, "llm-runtime-io:latest", 60*time.Second, "256m", 1)
	}

	relPath, err := filepath.Rel(repoRoot, filePath)
	if err != nil {
		return fmt.Errorf("failed to get relative path: %w", err)
	}

	// Atomic write: write to temp file then move
	command := fmt.Sprintf("mkdir -p $(dirname /workspace/%s) && printf '%%s' %q > /workspace/%s.tmp && mv /workspace/%s.tmp /workspace/%s",
		relPath, content, relPath, relPath, relPath)

	_, err = ExecuteInPooledContainer(ctx, pool, command, repoRoot)
	return err
}
