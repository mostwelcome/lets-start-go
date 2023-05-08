package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var employees = []Employee{
	{ID: 1, Name: "Alice", Role: "Software Engineer"},
	{ID: 2, Name: "Bob", Role: "Project Manager"},
	{ID: 3, Name: "Charlie", Role: "QA Engineer"},
}

func main() {
	app := fiber.New()

	// Add logger middleware
	app.Use(logger.New())

	// Define a protected route
	app.Get("/protected", authenticationMiddleware, func(c *fiber.Ctx) error {
		return c.SendString("You have access to the protected endpoint.")
	})

	// Define a public route for fetching employee details
	app.Get("/employees/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		employee := getEmployeeByID(id)

		if employee != nil {
			return c.JSON(employee)
		}
		return c.Status(fiber.StatusNotFound).SendString("Employee not found")
	})

	// Start the server
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func authenticationMiddleware(c *fiber.Ctx) error {
	// Replace this with your actual authentication logic
	const apiKey = "my-api-key"

	// Get the API key from the request header
	requestApiKey := c.Get("x-api-key")

	// Check if the provided API key matches the expected value
	if requestApiKey != apiKey {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	return c.Next()
}

func getEmployeeByID(id string) *Employee {
	for _, employee := range employees {
		if fmt.Sprint(employee.ID) == id {
			return &employee
		}
	}
	return nil
}
