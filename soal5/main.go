package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"soal5/lib"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type File struct {
	ID        int    `json:"id"`
	Filename  string `json:"filename"`
	URL       string `json:"url"`
	UploadedAt string `json:"uploaded_at"`
}

const (
	maxFileSize = 10 * 1024 * 1024 
	rootPath    = "./uploads/"
)

var db *pgx.Conn

func init() {
	var err error
	db, err = lib.DB() 
	if err != nil {
		panic(fmt.Errorf("unable to connect to database: %v\n", err))
	}

	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		os.MkdirAll(rootPath, os.ModePerm)
	}
}


func SaveFileMetadata(filename, url string) error {
	sql := `INSERT INTO files (filename, url) VALUES ($1, $2) RETURNING id`
	var id int
	err := db.QueryRow(context.Background(), sql, filename, url).Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to save file metadata: %w", err)
	}
	return nil
}


func UploadFile(c *gin.Context) {

	c.Request.ParseMultipartForm(maxFileSize)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no file is received"})
		return
	}


	if file.Size > maxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is too large. Maximum size is 10 MB."})
		return
	}

	allowExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".pdf": true, ".txt": true}
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	if !allowExt[fileExt] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file extension"})
		return
	}


	filename := uuid.New().String() + fileExt
	filePath := filepath.Join(rootPath, filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save the file"})
		return
	}


	fileURL := fmt.Sprintf("http://localhost:8000/files/%s", filename)


	if err := SaveFileMetadata(filename, fileURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": fileURL})
}


func GetFile(c *gin.Context) {
	filename := c.Param("filename")
	filePath := filepath.Join(rootPath, filename)


	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}

	c.File(filePath)
}

func SetupRouter() *gin.Engine {
	r := gin.Default()


	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	r.POST("/upload", UploadFile)
	r.GET("/files/:filename", GetFile)

	return r
}

func main() {
	defer db.Close(context.Background()) 

	r := SetupRouter()


	if err := r.Run("0.0.0.0:8000"); err != nil {
		fmt.Printf("failed to start server: %v\n", err)
	}
}
