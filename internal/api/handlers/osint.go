package handlers

import (
	"net/http"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"ctf-toolkit-backend/internal/utils"
)

// SherlockSearch handles the OSINT tool for username search using the Sherlock script.
func SherlockSearch(c *fiber.Ctx) error {
	// Parse the username from the request body
	type requestBody struct {
		Username string `json:"username" validate:"required"`
	}
	var body requestBody
	if err := c.BodyParser(&body); err != nil {
		return utils.JSONResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	// Execute the Sherlock script
	cmd := exec.Command("python3", "scripts/osint/sherlock_search.py", body.Username)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return utils.JSONResponse(c, http.StatusInternalServerError, "Error executing script: "+err.Error())
	}

	// Return the output as a JSON response
	return utils.JSONResponse(c, http.StatusOK, string(output))
}