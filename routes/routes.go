package routes

import (
    "github.com/grupoG/csw24-grupoG-ticket-gin/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter(sampleController *controllers.SampleController) *gin.Engine {
    router := gin.Default()

    api := router.Group("/api")
    {
        api.GET("/samples", sampleController.GetAllSamples)
        api.POST("/samples", sampleController.CreateSample)
    }

    return router
}
