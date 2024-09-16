package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func loadEnv(filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
            continue
        }
        parts := strings.SplitN(line, "=", 2)
        if len(parts) != 2 {
            continue
        }
        key := parts[0]
        value := parts[1]
        fmt.Printf("export %s=%s\n", key, value)
    }

    return scanner.Err()
}

func main() {
    if err := loadEnv("../.env"); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}