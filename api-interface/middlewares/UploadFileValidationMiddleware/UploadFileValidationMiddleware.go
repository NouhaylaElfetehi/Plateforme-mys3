package Middlewares

import (
	entities "api-interface/entities"
	"encoding/xml"
	"log"

	"github.com/gofiber/fiber/v2"
)

// UploadFileValidationMiddleware valide le corps de la requête pour l'upload de fichiers
func UploadFileValidationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var uploadRequest entities.UploadFileRequest
		if err := xml.Unmarshal(c.Body(), &uploadRequest); err != nil {
			log.Printf("Erreur XML: %v", err)
			return c.Status(fiber.StatusBadRequest).SendString("Le format XML est invalide: " + err.Error())
		}

		if uploadRequest.FileName == "" || uploadRequest.FileContent == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Le nom du fichier et son contenu sont requis")
		}

		// Stocker la requête dans Locals pour l'utiliser dans le contrôleur
		c.Locals("uploadRequest", &uploadRequest)

		// Passer au contrôleur
		return c.Next()
	}
}
