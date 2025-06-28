package handlers

import (
	"encoding/json"
	"log"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"ctf-toolkit-backend/internal/utils"
)

type ExifResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func ExtractExif(c *fiber.Ctx) error {
	// Get uploaded file
	file, err := c.FormFile("image")
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "No image file provided")
	}

	// Save uploaded file
	filename := filepath.Join("uploads", time.Now().Format("20060102150405")+"_"+file.Filename)
	if err := c.SaveFile(file, filename); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to save uploaded file")
	}

	// Execute EXIF extraction script
	cmd := exec.Command("python3", "scripts/stego/exif_extractor.py", filename)
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error executing EXIF extraction: %v", err)
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to extract EXIF data")
	}

	var result map[string]interface{}
	if err := json.Unmarshal(output, &result); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to parse result")
	}

	return utils.SuccessResponse(c, result, "EXIF data extracted successfully")
}