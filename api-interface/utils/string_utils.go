package utils

import "unicode"

// IsPascalCase vérifie si une chaîne est en PascalCase
func IsPascalCase(s string) bool {
    if len(s) == 0 {
        return false
    }
    if !unicode.IsUpper(rune(s[0])) {
        return false
    }
    for _, r := range s {
        if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
            return false
        }
    }
    return true
}

