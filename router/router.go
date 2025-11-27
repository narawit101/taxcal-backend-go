package router

import (
    "test-backend/handlers"

    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/tax/calculations", handlers.CalculateTaxHandler)
    r.POST("/tax/calculations2", handlers.CalculateTaxHandler2)

    return r
}
