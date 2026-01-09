package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "backend/database"
    "backend/models"
    "backend/services"
)

func CreateOrder(c *gin.Context) {
    var req struct {
        ItineraryID uint `json:"itineraryId"`
    }
    c.BindJSON(&req)

    var it models.Itinerary
    database.DB.First(&it, req.ItineraryID)

    data := map[string]interface{}{
        "amount": it.Price * 100,
        "currency": "INR",
    }

    order, _ := services.Client.Order.Create(data, nil)

    dbOrder := models.Order{
        ItineraryID: it.ID,
        Amount: it.Price,
        RazorpayOrder: order["id"].(string),
    }
    database.DB.Create(&dbOrder)

    c.JSON(http.StatusOK, gin.H{"orderId": order["id"]})
}
