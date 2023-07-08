package main

import (
	"log"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Serve static files (HTML, CSS, JS)
	router.Static("/static", "./static")

	// Define the API routes
	router.POST("/upload", uploadFile)
	router.GET("/download/:hash", downloadFile)

	// Set up the index route
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Start the server
	log.Println("Server started on http://localhost:8888")
	log.Fatal(router.Run(":8888"))
}

func uploadFile(c *gin.Context) {
	// Connect to the IPFS Cluster API
	sh := shell.NewShell("localhost:5001")

	// Get the uploaded file from the request
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	// Add the file to IPFS
	hash, err := sh.Add(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hash": hash})
}

func downloadFile(c *gin.Context) {
	// Connect to the IPFS Cluster API
	sh := shell.NewShell("localhost:5001")

	// Get the hash parameter from the URL
	hash := c.Param("hash")

	// Download the file from IPFS
	dstPath := filepath.Join("./downloads", hash)
	err := sh.Get(hash, dstPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set the appropriate content type for the response
	ext := filepath.Ext(dstPath)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// Serve the file as a response
	c.File(dstPath)
}
