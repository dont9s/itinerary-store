package main

import (
	"backend/database"
	"backend/models"
	"backend/routes"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(
		&models.Itinerary{},
		&models.Order{},
	)
	services.InitRazorpay()

	database.DB.Create(&models.Itinerary{
		Title:       "Goa Weekend Trip",
		Description: "3-day Goa beach itinerary",
		ImageURL:    "https://picsum.photos/600",
		Duration:    "3 Days",
		Price:       999,
		FilePath:    "goa.pdf",
	})

	r := gin.Default()
	routes.Register(r)
	r.Run(":8080")

}
