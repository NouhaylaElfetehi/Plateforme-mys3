package controllers

import (
	entities "api-interface/entities"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (b *BucketController) ListFiles(c *fiber.Ctx) error {
	bucketName := c.Params("bucketName")
	if bucketName == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Le nom du bucket est requis")
	}

	// Vérifier si le bucket existe
	bucketPath := filepath.Join("./buckets", bucketName)
	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).SendString("Le bucket n'existe pas")
	}

	var files []entities.BucketObject

	// Parcourir les fichiers du bucket
	err := filepath.Walk(bucketPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			relativePath, _ := filepath.Rel(bucketPath, path)
			files = append(files, entities.BucketObject{
				Key:          relativePath,
				LastModified: info.ModTime().Format(time.RFC3339),
				Size:         info.Size(),
				URI:          "http://localhost:3000/buckets/" + bucketName + "/" + relativePath,
				Type:         "OBJECT",
				BucketName:   bucketName,
			})
		}
		return nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la récupération des fichiers: " + err.Error())
	}

	// Renvoyer la liste des fichiers en XML
	return c.XML(files)
}
