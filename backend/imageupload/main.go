package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static files from the "uploads" directory
	r.Static("/uploads", "./uploads")

	// Set up CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           time.Hour,
	}))

	// Endpoint to handle file uploads
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		ImageOne, errone := c.FormFile("ImageOne")
		ImageTwo, errtwo := c.FormFile("ImageTwo")
		ImageThree, errthree := c.FormFile("ImageThree")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		if errone != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", errone.Error()))
			return
		}
		if errtwo != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", errtwo.Error()))
			return
		}
		if errthree != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", errthree.Error()))
			return
		}

		// Create the uploads directory if it doesn't exist
		_ = c.SaveUploadedFile(file, filepath.Join("uploads", file.Filename))
		_ = c.SaveUploadedFile(ImageOne, filepath.Join("uploads", ImageOne.Filename))
		_ = c.SaveUploadedFile(ImageTwo, filepath.Join("uploads", ImageTwo.Filename))
		_ = c.SaveUploadedFile(ImageThree, filepath.Join("uploads", ImageThree.Filename))

		c.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"url":      fmt.Sprintf("%s/uploads/%s", "http://localhost:3041", file.Filename),
			"urlone":   fmt.Sprintf("%s/uploads/%s", "http://localhost:3041", ImageOne.Filename),
			"urltwo":   fmt.Sprintf("%s/uploads/%s", "http://localhost:3041", ImageTwo.Filename),
			"urlthree": fmt.Sprintf("%s/uploads/%s", "http://localhost:3041", ImageThree.Filename),
		})
	})

	r.Run(":3041")
}
