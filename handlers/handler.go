package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/TFMV/VaultDeploy/models"
	"github.com/TFMV/VaultDeploy/services"
	"github.com/TFMV/VaultDeploy/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

func UploadCSVHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	filePath := filepath.Join(".", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	records, err := services.ProcessCSV(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process CSV"})
		return
	}

	config, err := utils.LoadConfig("config.yaml")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load config"})
		return
	}

	dbpool, err := utils.ConnectDatabase(config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer dbpool.Close()

	schema := config.DataVault.Schema

	for _, hubConfig := range config.DataVault.Hubs {
		hub := models.Hub{
			TableName: hubConfig.Name,
			Columns:   hubConfig.Columns,
		}
		err = services.CreateOrUpdateHub(dbpool, schema, hub)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create or update hub"})
			return
		}
	}

	for _, linkConfig := range config.DataVault.Links {
		link := models.Link{
			TableName: linkConfig.Name,
			Columns:   linkConfig.Columns,
		}
		err = services.CreateOrUpdateLink(dbpool, schema, link)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create or update link"})
			return
		}
	}

	for _, satelliteConfig := range config.DataVault.Satellites {
		satellite := models.Satellite{
			TableName: satelliteConfig.Name,
			Columns:   satelliteConfig.Columns,
		}
		err = services.CreateOrUpdateSatellite(dbpool, schema, satellite)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create or update satellite"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "CSV processed successfully"})
}

func ConfigureHandler(c *gin.Context) {
	var config utils.Config

	if err := c.Bind(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid configuration"})
		return
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize configuration"})
		return
	}

	if err := os.WriteFile("config.yaml", data, 0644); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save configuration"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Configuration saved successfully"})
}
