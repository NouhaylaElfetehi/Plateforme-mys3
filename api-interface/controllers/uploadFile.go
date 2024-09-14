package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func (b *BucketController) UploadFile(c *fiber.Ctx) error {

	bucketName := c.Params("bucketName")

	if bucketName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Le nom du bucket est requis",
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Aucun fichier n'a été téléchargé",
		})
	}

	bucketPath := filepath.Join("./buckets", bucketName)

	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Le bucket n'existe pas",
		})
	}

	filePath := filepath.Join(bucketPath, file.Filename)

	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erreur lors de l'enregistrement du fichier: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Le fichier '%s' a été téléchargé avec succès dans le bucket '%s'", file.Filename, bucketName),
	})
}
