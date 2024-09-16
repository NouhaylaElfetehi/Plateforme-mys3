package controllers

import (
	entities "api-interface/entities"
	"api-interface/handlers/errors"
	"api-interface/models"
	utils "api-interface/utils"
	bucketUtils "api-interface/utils/bucket"

	"time"

	"github.com/gofiber/fiber/v2"
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
    bucketName := c.Locals("bucketName").(string)

    var bucket *entities.Bucket
    var owner entities.Owner = entities.Owner{
        UserKey:     "000",
        DisplayName: "default",
        Type:       "default",
        URI:        "default",
        ROLE: entities.Role{
            ID:   "0",
            Name: "default",
            Type: "default",
        },
        SecretKey: "000",
    }

    bucketPath, err := bucketUtils.CreateBucketDirectory(bucketName)
    if err != nil {
        return errors.HandleError(c, errors.ErrInternalServerError, "Erreur lors de la création du répertoire du bucket")
    }

    // Création de l'entité Bucket

    //nous créons le svaleurs par defaut
    bucket = &entities.Bucket{
        Name:         bucketName,
        CreationDate: utils.StringPointer(time.Now().String()),
        Owner:        owner,
        URI:          bucketUtils.GenerateBucketURI(bucketName),
        Type:         "PUBLIC",
        Versioning:   "default",
    }

    // Insérer le bucket dans la base de données
    if err = b.bucketService.Insert(bucket); err != nil {
        return errors.HandleError(c, errors.ErrInternalServerError, "Erreur lors de l'insertion du bucket dans la base de données")
    }
    
    BASE_URL := c.BaseURL()
    return c.JSON(fiber.Map{
        "Location": BASE_URL+"/"+bucketPath,
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
