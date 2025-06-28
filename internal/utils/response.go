package utils

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// JSONResponse sends a JSON response with the given status code and data.
func JSONResponse(c *fiber.Ctx, status int, messageOrData interface{}, data ...interface{}) error {
	response := Response{
		Success: status >= 200 && status < 300,
	}

	if len(data) > 0 {
		response.Message = messageOrData.(string)
		response.Data = data[0]
	} else {
		if response.Success {
			response.Data = messageOrData
		} else {
			response.Error = messageOrData.(string)
		}
	}

	return c.Status(status).JSON(response)
}

// SuccessResponse sends a successful JSON response
func SuccessResponse(c *fiber.Ctx, data interface{}, message string) error {
	return c.JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends an error JSON response
func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(Response{
		Success: false,
		Error:   message,
	})
}