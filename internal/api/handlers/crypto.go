package handlers

import (
	"encoding/json"
	"log"
	"os/exec"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"ctf-toolkit-backend/internal/utils"
)

type CaesarRequest struct {
	Text string `json:"text"`
	Key  int    `json:"key"`
}

func CaesarCipher(c *fiber.Ctx) error {
	var req CaesarRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Execute Caesar cipher script
	cmd := exec.Command("python3", "scripts/crypto/caesar.py", req.Text, strconv.Itoa(req.Key), "encrypt")
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error executing Caesar cipher: %v", err)
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to execute cipher")
	}

	var result map[string]interface{}
	if err := json.Unmarshal(output, &result); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to parse result")
	}

	return utils.SuccessResponse(c, result, "Caesar cipher executed successfully")
}