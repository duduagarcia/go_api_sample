package main

import (
    "github.com/grupoG/csw24-grupoG-ticket-gin/configs"
    "github.com/grupoG/csw24-grupoG-ticket-gin/controllers"
    "github.com/grupoG/csw24-grupoG-ticket-gin/repositories"
    "github.com/grupoG/csw24-grupoG-ticket-gin/routes"
    "github.com/grupoG/csw24-grupoG-ticket-gin/services"
)

func main() {
    // Configurar o banco de dados
    db := configs.SetupDatabase()

    // Inicializar camadas
    sampleRepository := repositories.NewSampleRepository(db)
    sampleService := services.NewSampleService(sampleRepository)
    sampleController := controllers.NewSampleController(sampleService)

    // Configurar as rotas
    router := routes.SetupRouter(sampleController)

    // Iniciar o servidor
    router.Run(":8080")
}
