package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "path/filepath"
    "backend/database"
    "backend/models"
)

func Download(c *gin.Context) {
    token := c.Param("token")

    var order models.Order
    if err := database.DB.Where("download_token = ? AND paid = 1", token).First(&order).Error; err != nil {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    var it models.Itinerary
    database.DB.First(&it, order.ItineraryID)

    c.FileAttachment(filepath.Join("files", it.FilePath), "itinerary.pdf")
}
