package routes

import (
	"api-interface/database"
	"api-interface/handlers"
	Controller "api-interface/controllers"

	"github.com/gofiber/fiber/v2"
)

// Router configure les routes de l'application
func Router(app *fiber.App) {

	BucketObjectController := Controller.NewBucketObjectController()

	// if errorBucketObjectController != nil {
	// 	fmt.Println(errorBucketObjectController)
	// }

	// Routes publiques
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Bienvenue sur la plateforme MyS3")
	// })
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("static/public/index.html")
	})

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Get("/buckets/:bucketName/:filename", BucketObjectController.DownloadFile)


	// Routes protégées
	protected := app.Group("")
	protected.Use(handlers.AuthRequired)
	protected.Post("/bucket", database.CreateBucket)
	protected.Get("/bucket/:bucketName/files", database.ListFiles)
	protected.Post("/bucket/:bucketName/upload", database.UploadFile)
	protected.Get("/bucket/:bucketName/file/:fileName", database.DownloadFile)
	protected.Delete("/bucket/:bucketName/file/:fileName", database.DeleteFile)

}
