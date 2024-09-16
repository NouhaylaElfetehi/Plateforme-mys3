package models

import (
	"fmt"
	"reflect"
	"strings"

	entity "api-interface/entities"
	"api-interface/utils"
	colors "api-interface/utils/colorPrint"
)

var (
    Blue = colors.Blue
    Green = colors.Green
    Yellow = colors.Yellow
    Red = colors.Red
    Grey = colors.Grey
    Reset = colors.Reset
)

/** Tests d'application d'un TypeEntity avant l'intégration d'un Repository
 * @param typeName string
 * @param structType reflect.Type
 * @return bool
 */
func AssureEntityCompliance(typeName string, structType reflect.Type) bool {

    fmt.Printf("\n%sTraitement de l'entité : %s%s\n", Blue, typeName, Reset)
    fmt.Println(strings.Repeat("-", 50))

    // Vérifier que le type implémente l'interface Entity
    if !entity.ImplementsEntity(structType) {
        fmt.Printf("%sErreur : L'entité %s n'implémente pas l'interface entity.Entity%s\n", Red, structType.Name(), Reset)
        return false
    }

    // Vérifier que le nom de l'entité est en PascalCase
    if !utils.IsPascalCase(structType.Name()) {
        fmt.Printf("%sErreur : Le nom de l'entité %s n'est pas en PascalCase%s\n", Red, structType.Name(), Reset)
        return false
    }

    // Afficher la structure complète du type
    fmt.Printf("Structure de %s :\n", structType.Name())
    fmt.Println(Grey + strings.Repeat("-", 50) + Reset)
    for i := 0; i < structType.NumField(); i++ {
        field := structType.Field(i)
        fmt.Printf("  %s%s%s %s%s%s `json:\"%s\"`\n", Blue, field.Name, Reset, Green, field.Type, Reset, field.Tag.Get("json"))
    }

    return true
}