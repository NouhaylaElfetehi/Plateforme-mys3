package controllers

import (
	utils "api-interface/utils"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func (b *BucketController) ListObjects(c *fiber.Ctx) error {
	bucketName := c.Params("bucketName")
	if bucketName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Le nom du bucket est requis",
		})
	}

	bucketPath := filepath.Join("./buckets", bucketName)

	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Le bucket n'existe pas",
		})
	}

	files := []string{}
	err := filepath.Walk(bucketPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			// Lire le fichier chiffré
			encryptedFileData, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Déchiffrer les données
			key := utils.GetEncryptionKey()
			decryptedData, err := utils.Decrypt(encryptedFileData, key)
			if err != nil {
				return err
			}

			relativePath, relErr := filepath.Rel(bucketPath, path)
			if relErr != nil {
				return relErr
			}
			files = append(files, relativePath+" : "+string(decryptedData))
		}
		return nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erreur lors de la récupération des fichiers: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"bucket":  bucketName,
		"objects": files,
	})
}
