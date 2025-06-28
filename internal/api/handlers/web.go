package handlers

import (
    "encoding/json"
    "log"
    "os/exec"

    "github.com/gofiber/fiber/v2"
    "ctf-toolkit-backend/internal/utils"
)

type HTTPRequest struct {
    URL    string `json:"url" validate:"required"`
    Method string `json:"method"`
    Data   string `json:"data"`
}

// HTTPRequestSimulator handles the HTTP request simulation
func HTTPRequestSimulator(c *fiber.Ctx) error {
    var req HTTPRequest
    if err := c.BodyParser(&req); err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
    }

    // Set default method
    if req.Method == "" {
        req.Method = "GET"
    }

    // Prepare the command to execute the Python script
    var cmd *exec.Cmd
    if req.Data != "" {
        cmd = exec.Command("python3", "scripts/web/http_simulator.py", req.URL, req.Method, req.Data)
    } else {
        cmd = exec.Command("python3", "scripts/web/http_simulator.py", req.URL, req.Method)
    }

    // Execute the command and capture the output
    output, err := cmd.Output()
    if err != nil {
        log.Printf("Error executing HTTP simulator: %v", err)
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Error executing HTTP request")
    }

    var result map[string]interface{}
    if err := json.Unmarshal(output, &result); err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to parse simulation result")
    }

    // Return the output as a JSON response
    return utils.SuccessResponse(c, result, "HTTP request simulated successfully")
}