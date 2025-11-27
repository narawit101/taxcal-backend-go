package router

import (
    "test-backend/handlers"

)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/tax/calculations", handlers.CalculateTaxHandler)

    return r
}
