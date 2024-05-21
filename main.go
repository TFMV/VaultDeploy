package main

import (
	"log"
	"net/http"

	"github.com/TFMV/VaultDeploy/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", handlers.UploadCSVHandler)
	r.POST("/configure", handlers.ConfigureHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
