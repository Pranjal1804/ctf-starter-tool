package routes

import (
    "github.com/gofiber/fiber/v2"
    "ctf-toolkit-backend/internal/api/handlers"
)

func SetupRoutes(app *fiber.App) {
    // API version group
    api := app.Group("/api/v1")

    // Health check
    api.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status":  "healthy",
            "service": "ctf-toolkit-api",
        })
    })

    // Authentication routes
    auth := api.Group("/auth")
    auth.Post("/register", handlers.Register)
    auth.Post("/login", handlers.Login)
    auth.Get("/profile", handlers.GetProfile)

    // Cryptography routes
    crypto := api.Group("/crypto")
    crypto.Post("/caesar", handlers.CaesarCipher)

    // Steganography routes
    stego := api.Group("/stego")
    stego.Post("/exif", handlers.ExtractExif)

    // Binary analysis routes
    binary := api.Group("/binary")
    binary.Post("/strings", handlers.StringsExtractor)

    // Web exploitation routes
    web := api.Group("/web")
    web.Post("/http-simulator", handlers.HTTPRequestSimulator)

    // Network forensics routes
    network := api.Group("/network")
    network.Post("/pcap", handlers.PCAPAnalyzer)

    // OSINT routes
    osint := api.Group("/osint")
    osint.Post("/sherlock", handlers.SherlockSearch)

    // Miscellaneous routes
    misc := api.Group("/misc")
    misc.Post("/qr", handlers.QRCodeGenerator)
}