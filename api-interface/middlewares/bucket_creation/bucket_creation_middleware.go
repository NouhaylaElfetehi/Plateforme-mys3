package middlewares

import (
    "encoding/xml"
    entities "api-interface/entities"
    "api-interface/handlers/errors"
    bucket_name_validation "api-interface/middlewares/bucket_name_validation"
    "github.com/gofiber/fiber/v2"
    "os"
    "path/filepath"
    "strings"
)

// BucketValidationMiddleware est le middleware de validation pour les buckets
func BucketValidationMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Lire le corps de la requête
        bodyRequest := c.Body()

        // Nouvelle instance de CreateBucketRequestStruct
        bucketRequest := new(entities.CreateBucketRequestStruct)

        // Décoder le corps XML dans la struct bucketRequest
        if err := xml.Unmarshal(bodyRequest, bucketRequest); err != nil {
            return errors.HandleError(c, errors.ErrUnprocessableEntity, "Erreur de parsing XML: " + err.Error())
        }

        // Créez un validateur pour le nom du bucket
        validator := bucket_name_validation.NewBucketNameValidator()
        result := validator.Validate(bucketRequest.Name)

        if len(result) > 0 {
            errorMessage := strings.Join(result, "; ")
            return errors.HandleError(c, errors.ErrBadRequest, errorMessage)
        }

        // Valider les champs du bucket
        if err := validateBucketFields(bucketRequest, c); err != nil {
            return err
        }

        // Vérifier si le bucket existe déjà
        if bucketExists(bucketRequest.Name) {
            return errors.HandleError(c, errors.ErrBadRequest, "Le bucket avec ce nom existe déjà.")
        }

        // Attacher les données validées au contexte
        c.Locals("bucketRequest", bucketRequest)

        return c.Next()
    }
}

// validateBucketFields valide les champs du bucket
func validateBucketFields(bucket *entities.CreateBucketRequestStruct, c *fiber.Ctx) error {
    if bucket.Owner == (entities.Owner{}) {
        return errors.HandleError(c, errors.ErrBadRequest, "Le bucket doit posséder un propriétaire.")
    }
    if bucket.Type == "" {
        return errors.HandleError(c, errors.ErrBadRequest, "Le type de bucket doit être précisé.")
    }
    if bucket.Versioning == "" {
        return errors.HandleError(c, errors.ErrBadRequest, "Pas d'état de version pour le bucket.")
    }
    return nil
}

// bucketExists vérifie si un dossier de bucket existe déjà
func bucketExists(bucketName string) bool {
    bucketPath := filepath.Join("./buckets", bucketName)
    info, err := os.Stat(bucketPath)
    return !os.IsNotExist(err) && info.IsDir()
}
