package controllers

import (
    "api-interface/models"
    "github.com/gofiber/fiber/v2"
    "path/filepath"
    "os"
)

// FileController permet de gérer les requêtes liées aux fichiers
type BucketObjectController struct{
	bucketObjectService *models.BucketObjectModel
}

// NewFileController initialise un FileController
func NewBucketObjectController() *BucketObjectController {
    return &BucketObjectController{}
}

// DownloadFile gère le téléchargement d'un fichier depuis un bucket
func (fc *BucketObjectController) DownloadFile(c *fiber.Ctx) error {
    bucketName := c.Params("bucketName")
    fileName := c.Params("fileName")

    // Construire le chemin du fichier
    filePath := filepath.Join("./buckets", bucketName, fileName)

    // // Vérifier si le fichier existe
    // fileInfo, err := os.Stat(filePath)
    // if err != nil {
    //     if os.IsNotExist(err) {
    //         return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Fichier non trouvé"})
    //     }
    //     return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erreur lors de la vérification du fichier"})
    // }

    // Ouvrir le fichier
    file, err := os.Open(filePath)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erreur lors de l'ouverture du fichier"})
    }
    defer file.Close()

    // Définir les en-têtes de réponse
    c.Set("Content-Type", "application/octet-stream")
    c.Set("Content-Disposition", "attachment; filename="+fileName)
    // c.Set("Content-Length", fileInfo)
    c.Set("Cache-Control", "public, max-age=3600") // Cache pendant 1 heure

    // Servir le fichier
    return c.SendStream(file)
}