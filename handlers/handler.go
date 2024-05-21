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
	// Get the uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	// Save the uploaded file
	filePath := filepath.Join(".", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Process the CSV file
	records, err := services.ProcessCSV(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process CSV"})
		return
	}

	// Load the configuration
	config, err := utils.LoadConfig("config.yaml")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load config"})
		return
	}

	// Connect to the database
	dbpool, err := utils.ConnectDatabase(config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer dbpool.Close()

	schema := config.DataVault.Schema

	// Create or update Hubs
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

	// Create or update Links
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

	// Create or update Satellites
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

	// Insert records into the appropriate Data Vault tables
	for _, record := range records {
		// Example: Insert into the first hub's table
		if len(config.DataVault.Hubs) > 0 {
			hubConfig := config.DataVault.Hubs[0]
			err = services.InsertRecord(dbpool, schema, hubConfig.Name, hubConfig.Columns, record)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert record into hub"})
				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "CSV processed and records inserted successfully"})
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
