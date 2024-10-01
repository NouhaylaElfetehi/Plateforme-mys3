// main.go
package main

import (
	"api-interface/database"
	"api-interface/handlers"
	"api-interface/models"
	"api-interface/routes"
	"api-interface/utils/colorPrint"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"go.etcd.io/bbolt"
)

var (
	port = flag.String("port", ":9000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

var (
	SERVEUR_PORT string
	BOLT_PATH string
)

func init() {
	// Charger les variables d'environnement
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(colorPrint.RedP("Error loading .env file"))
	}

	envMap , envMapErr := godotenv.Read("../.env")
	if envMapErr != nil {
		fmt.Println(colorPrint.RedP("Error loading .env file"))
	} else if envMap == nil {
		fmt.Println(colorPrint.RedP("Error loading .env file"))
	}


	SERVEUR_PORT = envMap["S3_PORT"]
	if SERVEUR_PORT == "" {
		fmt.Println(colorPrint.YellowP("WARN | SERVEUR_PORT is empty"))
		SERVEUR_PORT = "9000"
		fmt.Println("SERVEUR_PORT is set to", colorPrint.GreenP(SERVEUR_PORT))
	}
	BOLT_PATH = envMap["DB_BOLT_PATH"]
	if BOLT_PATH == "" {
		fmt.Println(colorPrint.YellowP("WARN | BOLT_PATH is empty"))
		BOLT_PATH = "mydb.db"
		fmt.Println("BOLT_PATH is set to ", colorPrint.GreenP(BOLT_PATH))
	}
}

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	db, err := bbolt.Open("./data/metadata.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repositories
	models.InitRepositories(db)
	
	// Connect with database
	database.Connect()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Routes
	routes.Router(app)

	// Bind handlers
	app.Get("/users", handlers.UserList)
	app.Post("/users", handlers.UserCreate)

	// Setup static files
	app.Static("/", "./static/public")

	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port
	log.Fatal(app.Listen(":" + SERVEUR_PORT))
}