package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// GenerateBucketURI génère une URI pour un bucket donné
func GenerateBucketURI(bucketName string) string {
    // On récupère la base URL dynamiquement
    baseURL := os.Getenv("BASE_URL")
    println("BASE_URL: ", baseURL)
    return fmt.Sprintf("%s/%s", baseURL, bucketName)
}

// CreateBucketDirectory crée un répertoire pour un bucket donné
func CreateBucketDirectory(bucketName string) (string, error) {
    basePath := "./data/buckets" 
    bucketPath := filepath.Join(basePath, bucketName)

    if err := os.MkdirAll(bucketPath, os.ModePerm); err != nil {
        return "", fmt.Errorf("Erreur lors de la création du bucket : %v", err)
    }

    // ON modfiier les \\ en /
    bucketPath = filepath.ToSlash(bucketPath)

    return bucketPath, nil
}
