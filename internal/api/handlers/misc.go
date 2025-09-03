package handlers

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"ctf-toolkit-backend/internal/utils"
)

// QRCodeGenerator handles the QR code generation requests
func QRCodeGenerator(c *fiber.Ctx) error {
	// Get the text input from the request
	text := c.FormValue("text")
	if text == "" {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Text input is required")
	}

	// Generate unique filename
	timestamp := time.Now().Format("20060102150405")
	filename := filepath.Join("uploads", "qr_"+timestamp+".png")

	// Execute the QR code generator script with filename
	cmd := exec.Command("python3", "scripts/misc/qr_generator.py", text, filename)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Error generating QR code: "+err.Error())
	}

	var result map[string]interface{}
	if err := json.Unmarshal(output, &result); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to parse result")
	}

	// Add file path to response
	result["file_path"] = filename
	result["download_url"] = "/api/v1/files/" + filepath.Base(filename)

	return utils.SuccessResponse(c, result, "QR code generated and saved successfully")
}