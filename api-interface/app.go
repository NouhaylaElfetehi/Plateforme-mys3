// main.go
package main

import (
	"api-interface/database"
	"api-interface/handlers"
	"api-interface/models"
	"api-interface/routes"
	"api-interface/s3_operations"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"go.etcd.io/bbolt"
)

var (
	port     = flag.String("port", ":3000", "Port to listen on")
	prod     = flag.Bool("prod", false, "Enable prefork in Production")
	s3Client *s3.Client
)

func init() {
	// Charger les variables d'environnement
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Charger la configuration AWS
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-3"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	s3Client = s3.NewFromConfig(cfg)
}
func getEncryptionKey() []byte {
	// Récupérer la clé depuis l'environnement
	key := os.Getenv("ENCRYPTION_KEY")
	if len(key) != 32 {
		log.Fatalf("La clé de chiffrement doit avoir 32 octets pour AES-256")
	}
	return []byte(key)
}

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	db, err := bbolt.Open("my.db", 0600, nil)
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

	// Utiliser le serveur local comme base URL
	localBaseURL := "http://localhost:9000" // Exemple d'URL pour votre serveur local

	app.Get("/s3/objects", func(c *fiber.Ctx) error {
		lister := s3_operations.NewS3ObjectLister(s3Client, "your-access-key-id", "your-secret-access-key", localBaseURL)
		bucketName := c.Query("bucket")
		prefix := c.Query("prefix")
		delimiter := c.Query("delimiter")
		maxKeys := int32(10)

		if bucketName == "" {
			return c.Status(400).SendString("Bucket name is required")
		}

		objects, err := lister.ListObjects(context.TODO(), bucketName, prefix, delimiter, maxKeys)
		if err != nil {
			log.Printf("Error listing objects: %v", err)
			return c.Status(500).SendString("Failed to list objects")
		}

		return c.JSON(objects)
	})

	// Route pour supprimer un fichier
	app.Delete("/s3/object", func(c *fiber.Ctx) error {
		bucket := c.Query("bucket")
		key := c.Query("key")

		if bucket == "" || key == "" {
			return c.Status(400).SendString("Bucket name and key are required")
		}

		req := &s3.DeleteObjectInput{
			Bucket: &bucket,
			Key:    &key,
		}

		_, err := s3Client.DeleteObject(context.TODO(), req)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Failed to delete object: %v", err))
		}

		return c.SendString("Object deleted successfully")
	})

	// Setup static files
	app.Static("/", "./static/public")

	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port
	log.Fatal(app.Listen(*port))
}
