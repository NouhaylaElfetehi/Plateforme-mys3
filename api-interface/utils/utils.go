package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
	"os"
)

// Decrypt déchiffre les données chiffrées à l'aide d'une clé AES-256
func Decrypt(encryptedData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la création du cipher: %v", err)
	}

	if len(encryptedData) < aes.BlockSize {
		return nil, fmt.Errorf("données chiffrées trop courtes")
	}

	iv := encryptedData[:aes.BlockSize]
	encryptedData = encryptedData[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	decryptedData := make([]byte, len(encryptedData))
	stream.XORKeyStream(decryptedData, encryptedData)

	return decryptedData, nil
}

// GenerateRandomKey génère une clé de chiffrement aléatoire
func GetEncryptionKey() []byte {
	key := os.Getenv("ENCRYPTION_KEY")
	if len(key) != 32 {
		log.Fatalf("La clé de chiffrement doit être de 32 octets pour AES-256")
	}
	return []byte(key)
}
