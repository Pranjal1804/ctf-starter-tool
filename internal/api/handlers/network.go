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

// PCAPAnalyzer handles the PCAP file upload and analysis
func PCAPAnalyzer(c *fiber.Ctx) error {
    // Parse the uploaded file
    file, err := c.FormFile("pcap")
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "No PCAP file provided")
    }

    // Validate file extension
    ext := filepath.Ext(file.Filename)
    if ext != ".pcap" && ext != ".pcapng" && ext != ".cap" {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid file type. Please upload a PCAP file")
    }

    // Save the file temporarily
    tempFilename := filepath.Join("uploads", time.Now().Format("20060102150405")+"_"+file.Filename)
    if err := c.SaveFile(file, tempFilename); err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to save uploaded file")
    }

    // Clean up the temp file after analysis
    defer func() {
        if err := os.Remove(tempFilename); err != nil {
            log.Printf("Failed to remove temp file: %v", err)
        }
    }()

    // Execute the Python script for PCAP analysis
    cmd := exec.Command("python3", "scripts/network/pcap_analyzer.py", tempFilename)
    output, err := cmd.Output()
    if err != nil {
        log.Printf("Error executing PCAP analysis: %v", err)
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Error executing analysis")
    }

    var result map[string]interface{}
    if err := json.Unmarshal(output, &result); err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to parse analysis result")
    }

    // Return the analysis result
    return utils.SuccessResponse(c, result, "PCAP analysis completed successfully")
}