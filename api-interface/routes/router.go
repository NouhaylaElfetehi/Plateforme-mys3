package routes

import (
	Controller "api-interface/controllers"
	"api-interface/database"
	"api-interface/handlers"
<<<<<<< HEAD
	Controller "api-interface/controllers"
=======
	Middlewares "api-interface/middlewares/bucket_creation"

	"fmt"
>>>>>>> main

	"github.com/gofiber/fiber/v2"
)

// Router configure les routes de l'application
func Router(app *fiber.App) {

	BucketObjectController := Controller.NewBucketObjectController()

	// if errorBucketObjectController != nil {
	// 	fmt.Println(errorBucketObjectController)
	// }

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

	app.Get("/:bucketName/:filename", BucketObjectController.DownloadFile)


	// Routes protégées
	protected := app.Group("")
	protected.Use(handlers.AuthRequired)

	// Route de création du Bucket
	protected.Post("/bucket", database.CreateBucket)
	protected.Get("/bucket/:bucketName/files", database.ListFiles)
	protected.Post("/bucket/:bucketName/upload", database.UploadFile)
	protected.Get("/bucket/:bucketName/file/:fileName", database.DownloadFile)
	protected.Delete("/bucket/:bucketName/file/:fileName", database.DeleteFile)

<<<<<<< HEAD
=======
	// Create Bucket
	app.Put("/:bucketName", Middlewares.BucketValidationMiddleware(),  bc.InsertBucket)
>>>>>>> main
}
