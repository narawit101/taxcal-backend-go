package router

import (
    "test-backend/handlers"

    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/tax/calculations/simple", handlers.CalculateTaxHandlerSimple)
    r.POST("/tax/calculations/detail", handlers.CalculateTaxHandlerDetail)

    return r
}
