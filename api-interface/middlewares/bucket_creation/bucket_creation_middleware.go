package middlewares

import (
	dto "api-interface/entities/bucketdtos"
	"api-interface/handlers/errors"
	bucket_name_validation "api-interface/middlewares/bucket_name_validation"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// BucketValidationMiddleware est le middleware de validation pour les buckets
func BucketValidationMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Lire le corps de la requête
        bodyRequest := c.Body()
        bucketName := c.Params("bucketName")


        // Nouvelle instance de CreateBucketRequestConfiguration
        bucketRequest := new(dto.CreateBucketConfiguration)
        // Décoder le corps XML dans la struct bucketRequest
        if err := xml.Unmarshal(bodyRequest, bucketRequest); err != nil {
            return errors.HandleError(c, errors.ErrUnprocessableEntity, "Erreur de parsing XML: " + err.Error())
        }

        errFormatXML := validateBucketFields(bucketRequest, c)
        // Valider les champs du bucket
        if errFormatXML != "" {
            return errors.HandleError(c, errors.ErrBadRequest, errFormatXML)
            // ON renvoie une response d'erreur
        } else {
            fmt.Println("Validation des champs du bucket réussie.")
        }

        c.Locals("bucketRequest", bucketRequest)


        // Créez un validateur pour le nom du bucket
        validator := bucket_name_validation.NewBucketNameValidator()
        result := validator.Validate(bucketName)

        if len(result) > 0 {
            errorMessage := strings.Join(result, "; ")
            return errors.HandleError(c, errors.ErrBadRequest, errorMessage)
        }

        // Vérifier si le bucket existe déjà
        if bucketExists(bucketName) {
            return errors.HandleError(c, errors.ErrConflict, "Le bucket avec ce nom existe déjà.")
        }

        // Attacher les données validées au contexte
        c.Locals("bucketName", bucketName)

        return c.Next()
    }
}

// validateBucketFields valide les champs du bucket
func validateBucketFields(bucketConfig *dto.CreateBucketConfiguration, c *fiber.Ctx) string {
    bucketConfig.XMLName.Local = "CreateBucketConfiguration"
    bucketConfig.XMLNS = "http://s3.amazonaws.com/doc/2006-03-01/"

    // Valider le champ LocationConstraint
    if bucketConfig.LocationConstraint != nil {
        if !dto.IsValidLocationConstraint(*bucketConfig.LocationConstraint) {
            return "LocationConstraint invalide."
        }
    } else {
        fmt.Println("LocationConstraint non défini, défini sur 'us-east-1'")
        bucketConfig.LocationConstraint = new(string)
        *bucketConfig.LocationConstraint = "us-east-1"
    }

    // Valider le champ Location
    if bucketConfig.Location != nil {
        if !dto.IsValidLocationConstraint(bucketConfig.Location.Name) {
            return "Location invalide"
        }
        if !dto.IsValidLocationType(bucketConfig.Location.Type) {

            return "Location Type invalide, doit être de type 'AvailabilityZone'."
        }
    } else {
        bucketConfig.Location = &dto.LocationInfo{
            Name: "EU",
            Type: "AvailabilityZone",
        }
    }

    // Valider le champ Bucket
    if bucketConfig.Bucket != nil {
        if !dto.IsValidDataRedundancy(bucketConfig.Bucket.DataRedundancy) {
            return "DataRedundancy invalide doit être de type 'SingleAvailabilityZone'."
        }
        if !dto.IsValidBucketType(bucketConfig.Bucket.Type) {
            return"Bucket Type invalide, doit être de type 'Directory'."
        }
    } else {
        bucketConfig.Bucket = &dto.BucketInfo{
            DataRedundancy: "SingleAvailabilityZone",
            Type:           "Directory",
        }
    }

    return ""

}

// bucketExists vérifie si un dossier de bucket existe déjà
func bucketExists(bucketName string) bool {
    bucketPath := filepath.Join("./data/buckets", bucketName)
    info, err := os.Stat(bucketPath)
    fmt.Println("Bucket Path: ", info)
    return !os.IsNotExist(err) && info.IsDir()
}
