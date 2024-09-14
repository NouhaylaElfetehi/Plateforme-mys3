package controllers

import (
	entities "api-interface/entities"
	"encoding/base64"
	"encoding/xml"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func (b *BucketController) UploadFiles(c *fiber.Ctx) error {
	// Récupérer les paramètres via le middleware
	uploadRequest, ok := c.Locals("uploadRequest").(*entities.UploadFileRequest)
	if !ok {
		return fiber.NewError(fiber.StatusInternalServerError, "Erreur interne.")
	}

	bucketName := c.Params("bucketName")
	if bucketName == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Le nom du bucket est requis")
	}

	// Vérifier si le bucket existe
	bucketPath := filepath.Join("./buckets", bucketName)
	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).SendString("Le bucket n'existe pas")
	}

	// Obtenir le chemin du fichier et décoder le contenu en base64
	filePath := filepath.Join(bucketPath, uploadRequest.FileName)
	fileContent, err := base64.StdEncoding.DecodeString(uploadRequest.FileContent)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Erreur lors du décodage du contenu du fichier: " + err.Error())
	}

	// Sauvegarder le fichier
	if err = os.WriteFile(filePath, fileContent, 0644); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la sauvegarde du fichier: " + err.Error())
	}

	// Réponse XML
	response := struct {
		XMLName xml.Name `xml:"UploadResponse"`
		Message string   `xml:"message"`
		Path    string   `xml:"path"`
	}{
		Message: "Fichier téléchargé avec succès",
		Path:    filePath,
	}

	return c.XML(response)
}
