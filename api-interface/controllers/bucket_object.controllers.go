package controllers

import (
    "api-interface/models"
    "github.com/gofiber/fiber/v2"
    "path/filepath"
    "os"
    "fmt"
)

// FileController permet de gérer les requêtes liées aux fichiers
type BucketObjectController struct{
	bucketObjectService *models.BucketObjectModel
}

// NewFileController initialise un FileController
func NewBucketObjectController() *BucketObjectController {
    return &BucketObjectController{}
}

func (fc *BucketObjectController) DownloadFile(c *fiber.Ctx) error {
    bucketName := c.Params("bucketName")
    fileName := c.Params("fileName")

    // Construire le chemin du fichier
    filePath := filepath.Join("./buckets", bucketName, fileName)
    fmt.Println("Chemin du fichier :", filePath)

    // Vérifier si le fichier existe
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Fichier non trouvé"})
    }

    // Utiliser SendFile pour servir le fichier
    return c.SendFile(filePath)
}
