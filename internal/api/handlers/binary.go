package handlers

import (
    "encoding/json"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "time"

    "github.com/gofiber/fiber/v2"
    "ctf-toolkit-backend/internal/utils"
)

// StringsExtractor handles the extraction of strings from binary files.
func StringsExtractor(c *fiber.Ctx) error {
    // Parse the uploaded file
    file, err := c.FormFile("file")
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "No binary file provided")
    }

    // Save the file temporarily
    tempFilename := filepath.Join("uploads", time.Now().Format("20060102150405")+"_"+file.Filename)
    if err := c.SaveFile(file, tempFilename); err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to save uploaded file")
    }

    // Clean up the temp file
    defer func() {
        if err := os.Remove(tempFilename); err != nil {
            log.Printf("Failed to remove temp file: %v", err)
        }
    }()

    // Execute the Python script
    cmd := exec.Command("python3", "scripts/binary/strings_extractor.py", tempFilename)
    output, err := cmd.Output()
    if err != nil {
        log.Printf("Error executing strings extraction: %v", err)
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Error extracting strings")
    }

    var result map[string]interface{}
    if err := json.Unmarshal(output, &result); err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to parse extraction result")
    }

    // Return the output as JSON
    return utils.SuccessResponse(c, result, "Strings extracted successfully")
}