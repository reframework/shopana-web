package middlewares

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Custom middleware to log request processing time
func RequestTime(c *fiber.Ctx) error {
	start := time.Now()

	// Process the request
	err := c.Next()

	// Calculate the duration
	duration := time.Since(start)

	// Log the request details and duration
	log.Printf("Method: %s, Path: %s, Status: %d, Duration: %s\n",
		c.Method(), c.Path(), c.Response().StatusCode(), duration)

	return err
}
