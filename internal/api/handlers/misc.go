package handlers

import (
	"net/http"
	"os/exec"

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

	// Execute the QR code generator script
	cmd := exec.Command("python3", "scripts/misc/qr_generator.py", text)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Error generating QR code: "+err.Error())
	}

	// Return the generated QR code as a response
	return utils.SuccessResponse(c, string(output), "QR code generated successfully")
}