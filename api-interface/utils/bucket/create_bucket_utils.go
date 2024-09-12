package utils

import (
    "fmt"
    "os"
    "path/filepath"
)

// GenerateBucketURI génère une URI pour un bucket donné
func GenerateBucketURI(bucketName string) string {
    baseURL := "http://localhost:3000/buckets"
    return fmt.Sprintf("%s/%s", baseURL, bucketName)
}

// CreateBucketDirectory crée un répertoire pour un bucket donné
func CreateBucketDirectory(bucketName string) (string, error) {
    basePath := "./buckets" 
    bucketPath := filepath.Join(basePath, bucketName)

    if err := os.MkdirAll(bucketPath, os.ModePerm); err != nil {
        return "", fmt.Errorf("Erreur lors de la création du bucket : %v", err)
    }

    return bucketPath, nil
}
