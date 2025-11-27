package router

import (
    "test-backend/handlers"

    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/tax/calculations", handlers.CalculateTaxHandler)

    return r
}
