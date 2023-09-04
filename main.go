package main

import (
	"flag"
	"github.com/infytvcode/dsp/database"
	"github.com/infytvcode/dsp/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	database.Connect()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://imasdk.googleapis.com",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	app.Post("/openrtb", handlers.OpenRTBHandler)
	app.Get("/vast", handlers.VastHandler)
	app.Get("/vmap", handlers.VmapHandler)
	app.Get("/vpaid", handlers.VpaidHandler)
	app.Get("/aws_vast", handlers.VastHandler)
	app.Get("/tracking", handlers.Tracking)

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Get("/users", handlers.CustomerList)
	v1.Post("/users", handlers.CustomerCreate)

	// Setup static files
	app.Static("/", "./static/public")

	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(*port))
}
