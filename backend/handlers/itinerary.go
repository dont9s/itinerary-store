package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend/database"
    "backend/models"
)

func GetItineraries(c *gin.Context) {
    var data []models.Itinerary
    database.DB.Find(&data)
    c.JSON(http.StatusOK, data)
}
