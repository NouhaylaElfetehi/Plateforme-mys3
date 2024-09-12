package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

// Fonction pour supprimer un objet dans un bucket
func (b *BucketController) DeleteObject(c *fiber.Ctx) error {
	// Récupérer le nom du bucket et le nom de l'objet depuis les paramètres de l'URL
	bucketName := c.Params("bucketName")
	objectName := c.Params("objectName")

	if bucketName == "" || objectName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Le nom du bucket et celui de l'objet sont requis",
		})
	}

	// Définir le chemin complet de l'objet à supprimer
	objectPath := filepath.Join("./buckets", bucketName, objectName)

	// Vérifier si le fichier existe
	if _, err := os.Stat(objectPath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "L'objet n'existe pas dans le bucket",
		})
	}

	// Supprimer le fichier
	err := os.Remove(objectPath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erreur lors de la suppression de l'objet: " + err.Error(),
		})
	}

	// Retourner une réponse de succès
	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("L'objet '%s' a été supprimé avec succès du bucket '%s'", objectName, bucketName),
	})
}
