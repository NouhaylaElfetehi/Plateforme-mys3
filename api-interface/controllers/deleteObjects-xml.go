package controllers

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func (b *BucketController) DeleteFile(c *fiber.Ctx) error {
	bucketName := c.Params("bucketName")
	fileName := c.Params("fileName")

	if bucketName == "" || fileName == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Le nom du bucket et du fichier sont requis")
	}

	// Vérifier si le bucket existe
	bucketPath := filepath.Join("./buckets", bucketName)
	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).SendString("Le bucket n'existe pas")
	}

	// Obtenir le chemin du fichier
	filePath := filepath.Join(bucketPath, fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).SendString("Le fichier n'existe pas")
	}

	// Supprimer le fichier
	if err := os.Remove(filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la suppression du fichier")
	}

	return c.Status(fiber.StatusOK).SendString("Fichier supprimé avec succès")
}
