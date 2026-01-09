package routes

import (
    "github.com/gin-gonic/gin"
    "backend/handlers"
)

func Register(r *gin.Engine) {
    api := r.Group("/api")
    {
        api.GET("/itineraries", handlers.GetItineraries)
        api.POST("/create-order", handlers.CreateOrder)
        api.GET("/download/:token", handlers.Download)
    }
}
