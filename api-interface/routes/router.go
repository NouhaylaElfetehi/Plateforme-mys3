package routes

import (
	"api-interface/database"
	"api-interface/handlers"
	Controller "api-interface/controllers"
	Middlewares "api-interface/middlewares/bucket_creation"

	"github.com/gofiber/fiber/v2"
	"fmt"
)

// Router configure les routes de l'application
func Router(app *fiber.App) {

 	bc, errorBc := Controller.NewBucketController()

	if errorBc != nil {
		fmt.Println(errorBc)
	}
	// Routes publiques
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Bienvenue sur la plateforme MyS3")
	// })
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("static/public/index.html")
	})

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	// Routes protégées
	protected := app.Group("")
	protected.Use(handlers.AuthRequired)

	// Route de création du Bucket
	protected.Post("/bucket", database.CreateBucket)
	protected.Get("/bucket/:bucketName/files", database.ListFiles)
	protected.Post("/bucket/:bucketName/upload", database.UploadFile)
	protected.Get("/bucket/:bucketName/file/:fileName", database.DownloadFile)
	protected.Delete("/bucket/:bucketName/file/:fileName", database.DeleteFile)

	protected.Post("/v1/bucket", Middlewares.BucketValidationMiddleware(),  bc.InsertBucket)
}
