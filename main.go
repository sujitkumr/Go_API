// package main

// import (
// 	"github.com/gofiber/fiber/v2"
// )

// func main() {
// 	app := fiber.New()

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("Hello, Fiber!")
// 	})

// 	app.Listen(":3000")
// }

package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sujitkumr/go_api/config"
	"github.com/sujitkumr/go_api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	_ = godotenv.Load(".env") // Ensure to load .env file
}

func main() {
	// Load environment variables
	mongoURI := os.Getenv("DATABASE_URL")

	log.Println("DATABASE_URL:", os.Getenv("DATABASE_URL"))

	if mongoURI == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Initialize the database
	config.ConnectDB(mongoURI)

	// Create a Fiber app
	app := fiber.New()

	// Middleware
	app.Use(cors.New())

	// Register routes
	routes.RegisterRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
