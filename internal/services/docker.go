package services

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

// DockerService provides methods to manage Docker containers for executing scripts.
type DockerService struct{}

// NewDockerService creates a new instance of DockerService.
func NewDockerService() *DockerService {
	return &DockerService{}
}

// RunScript executes a Python script in a Docker container and returns the output.
func (ds *DockerService) RunScript(scriptPath string, args []string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Prepare the command to run the Docker container
	cmdArgs := []string{"run", "--rm", "-v", fmt.Sprintf("%s:/scripts", scriptPath), "python:3.9", "/scripts/" + scriptPath}
	cmdArgs = append(cmdArgs, args...)

	cmd := exec.CommandContext(ctx, "docker", cmdArgs...)
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error executing script: %s, stderr: %s", err.Error(), stderr.String())
	}

	return out.String(), nil
}