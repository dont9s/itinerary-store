package main

import (
    "github.com/gin-gonic/gin"
    "backend/database"
    "backend/routes"
    "backend/services"
)

func main() {
    database.Connect()
    services.InitRazorpay()

    r := gin.Default()
    routes.Register(r)
    r.Run(":8080")
}
