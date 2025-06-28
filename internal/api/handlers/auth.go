package handlers

import (
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"

    "ctf-toolkit-backend/internal/config"
    "ctf-toolkit-backend/internal/database/models"
    "ctf-toolkit-backend/internal/utils"
)

// UserCredentials
type UserCredentials struct {
    Username string `json:"username" validate:"required"`
    Password string `json:"password" validate:"required"`
}

// RegisterRequest
type RegisterRequest struct {
    Username string `json:"username" validate:"required,min=3,max=20"`
    Password string `json:"password" validate:"required,min=6"`
    Email    string `json:"email" validate:"required,email"`
}

// Register handles user registration
func Register(c *fiber.Ctx) error {
    var req RegisterRequest
    if err := c.BodyParser(&req); err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
    }

    //validation
    if len(req.Username) < 3 || len(req.Password) < 6 {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Username must be at least 3 characters and password at least 6 characters")
    }

    //password hash
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Could not hash password")
    }

    // User Object
    user := models.User{
        Username: req.Username,
        Password: string(hashedPassword),
        Email:    req.Email,
    }

    // Store in database
    if err := models.SaveUser(user); err != nil {
        if err.Error() == "username already exists" || err.Error() == "email already exists" {
            return utils.ErrorResponse(c, fiber.StatusConflict, err.Error())
        }
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Could not register user")
    }

    return utils.SuccessResponse(c, fiber.Map{
        "username": user.Username,
        "email":    user.Email,
    }, "User registered successfully")
}

// user login handler.
func Login(c *fiber.Ctx) error {
    var creds UserCredentials
    if err := c.BodyParser(&creds); err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
    }

    // Find user by username
    user, err := models.FindUserByUsername(creds.Username)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid username or password")
    }

    // Check password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
        return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid username or password")
    }

    // Config load
    cfg, err := config.LoadConfig()
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Server configuration error")
    }

    // JWT token generation.
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id":  user.ID.Hex(),
        "username": user.Username,
        "exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
        "iat":      time.Now().Unix(),
    })

    tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Could not generate token")
    }

    return utils.SuccessResponse(c, fiber.Map{
        "token": tokenString,
        "user": fiber.Map{
            "id":       user.ID.Hex(),
            "username": user.Username,
            "email":    user.Email,
        },
    }, "Login successful")
}

// GetProfile returns the user's profile information
func GetProfile(c *fiber.Ctx) error {
    // This would typically extract user info from JWT token
    // For now, we'll implement a basic version
    return utils.SuccessResponse(c, fiber.Map{
        "message": "Profile endpoint - implement JWT middleware first",
    }, "Profile retrieved")
}