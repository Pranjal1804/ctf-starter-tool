package services

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"path/filepath"
)

// ExecutePythonScript executes a Python script with the given arguments and returns the output.
func ExecutePythonScript(scriptPath string, args []string) (map[string]interface{}, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer

	// Prepare the command
	cmd := exec.Command("python3", append([]string{scriptPath}, args...)...)

	// Set the output and error buffers
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Execute the command
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	// Parse the output as JSON
	var result map[string]interface{}
	if err := json.Unmarshal(out.Bytes(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetExifData extracts EXIF data from the provided image file.
func GetExifData(imagePath string) (map[string]interface{}, error) {
	scriptPath := filepath.Join("scripts", "stego", "exif_extractor.py")
	return ExecutePythonScript(scriptPath, []string{imagePath})
}