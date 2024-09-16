package errors

import (
	"github.com/gofiber/fiber/v2"
)

// Structure pour une erreur personnalisée
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Nouvelle instance d'erreur
func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// HandleError écrit la réponse d'erreur en format JSON dans le writer HTTP.
func HandleError(c *fiber.Ctx, err *Error, customMessage ...string) error {

	// Remplacer le message par défaut s'il y a un message personnalisé
	if len(customMessage) > 0 {
		err.Message = customMessage[0]
	}

	// Sérialiser l'erreur en JSON et la renvoyer comme réponse Fiber
	return c.Status(err.Code).JSON(fiber.Map{
		"error": err.Message,
	})
}

// Exemples d'utilisation avec message par défaut pour chaque statut.
var (
	// 400
	ErrBadRequest          = New(fiber.StatusBadRequest, "Bad Request")
	// 403
	ErrForbidden           = New(fiber.StatusForbidden, "Forbidden")
	// 404
	ErrNotFound            = New(fiber.StatusNotFound, "Not Found")
	// 409
	ErrConflict            = New(fiber.StatusConflict, "Conflict")
	// 413
	ErrRequestEntityTooLarge = New(fiber.StatusRequestEntityTooLarge, "Request entity too large")
	// 422
	ErrUnprocessableEntity = New(fiber.StatusUnprocessableEntity, "Unprocessable Entity") 
	// 500
	ErrInternalServerError = New(fiber.StatusInternalServerError, "Internal Server Error")
)