package main

import (
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

var (
    fileNames []string
    mutex     sync.RWMutex
)

func init() {
    loadFileNames()
}

func loadFileNames() {
    files, err := filepath.Glob("./data/*")
    if err != nil {
        log.Fatalf("Failed to list files: %v", err)
    }

    fileNames = make([]string, len(files))
    for i, file := range files {
        fileNames[i] = filepath.Base(file)
    }

    sort.Strings(fileNames)
}

func binarySearch(name string) string {
    low, high := 0, len(fileNames)-1
    for low <= high {
        mid := (low + high) / 2
        if fileNames[mid] == name {
            return fileNames[mid]
        } else if fileNames[mid] < name {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return ""
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Failed to upload file", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    filePath := "./data/" + header.Filename
    out, err := os.Create(filePath)
    if err != nil {
        http.Error(w, "Failed to save file", http.StatusInternalServerError)
        return
    }
    defer out.Close()

    _, err = io.Copy(out, file)
    if err != nil {
        http.Error(w, "Failed to save file", http.StatusInternalServerError)
        return
    }

    mutex.Lock()
    fileNames = append(fileNames, header.Filename)
    sort.Strings(fileNames)
    mutex.Unlock()

    w.WriteHeader(http.StatusCreated)
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")

    mutex.RLock()
    fileName := binarySearch(name)
    mutex.RUnlock()

    if fileName == "" {
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }

    filePath := "./data/" + fileName
    file, err := os.Open(filePath)
    if err != nil {
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }
    defer file.Close()

    fileType := mime.TypeByExtension(filepath.Ext(fileName))
    if fileType == "" {
        fileType = "application/octet-stream"
    }

    w.Header().Set("Content-Type", fileType)
    http.ServeFile(w, r, filePath)
}

func main() {
    http.HandleFunc("/upload", uploadFile)
    http.HandleFunc("/download", downloadFile)

    log.Fatal(http.ListenAndServe(":8081", nil))
}