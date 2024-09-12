package controllers

import (
    utils "api-interface/utils"
    bucketUtils "api-interface/utils/bucket"
    entities "api-interface/entities"
    "api-interface/models"
    "api-interface/handlers/errors"

    "github.com/gofiber/fiber/v2"
    "time"
)

/** Bucket Controller permet de gérer les requêtes liées aux Buckets */
type BucketController struct {
    bucketService *models.BucketModel
}

/** NewBucketController initialise un BucketController avec le service approprié */
func NewBucketController() (*BucketController, error) {
    bucketService, err := models.UseBucketModel()
    if err != nil {
        return nil, err
    }

    return &BucketController{
        bucketService: bucketService,
    }, nil
}

// InsertBucket gère l'insertion d'un nouveau bucket.
// Il utilise un middleware de validation des données, ainsi que des utils permettant
// De générer l'URI du bucket et créer le dossier sur le serveur.
func (b *BucketController) InsertBucket(c *fiber.Ctx) error {

    // Utiliser le middleware pour la validation des buckets
    bucketRequest, ok := c.Locals("bucketRequest").(*entities.CreateBucketRequestStruct)
    if !ok {
        return fiber.NewError(fiber.StatusInternalServerError, "Erreur interne.")
    }

    // Générer l'URI et créer le répertoire pour le bucket
    bucketURI := bucketUtils.GenerateBucketURI(bucketRequest.Name)
    bucketPath, err := bucketUtils.CreateBucketDirectory(bucketRequest.Name)
    if err != nil {
        return errors.HandleError(c, errors.ErrInternalServerError, "Erreur lors de la création du répertoire du bucket")
    }

    defaultBucketCreationDate := time.Now()

    // Création de l'entité Bucket
    bucket := &entities.Bucket{
        Name:         bucketRequest.Name,
        CreationDate: utils.StringPointer(defaultBucketCreationDate.String()),
        Owner:        bucketRequest.Owner,
        URI:          bucketURI,
        Type:         bucketRequest.Type,
        Versioning:   bucketRequest.Versioning,
    }

    // Insérer le bucket dans la base de données
    if err = b.bucketService.Insert(bucket); err != nil {
        return errors.HandleError(c, errors.ErrInternalServerError, "Erreur lors de l'insertion du bucket dans la base de données")
    }

    return c.JSON(fiber.Map{
        "message": "Bucket inséré avec succès",
        "uri":     bucketURI,
        "path":    bucketPath,
    })
}

//* getAllBuckets permet de récupérer tous les Buckets de la base de données
func (b *BucketController) GetAllBuckets(c *fiber.Ctx) error {
    buckets, err := b.bucketService.GetAllBuckets()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(buckets)
}

/** GetBucketByName permet de récupérer un Bucket par son nom */
func (b *BucketController) GetBucketByName(c *fiber.Ctx) error {
    name := c.Params("name")
    bucket, err := b.bucketService.GetBucketByName(name)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(bucket)
}
