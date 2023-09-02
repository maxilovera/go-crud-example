package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maxilovera/go-crud-example/handlers"
	"github.com/maxilovera/go-crud-example/repositories"
	"github.com/maxilovera/go-crud-example/services"
)

var (
	aulaHandler *handlers.AulaHandler
	router      *gin.Engine
)

func main() {
	router = gin.Default()
	//Iniciar objetos de handler
	dependencies()
	//Iniciar rutas
	mappingRoutes()

	router.Run(":8080")
}

func mappingRoutes() {
	//Listado de rutas
	router.GET("/aulas", aulaHandler.ObtenerAulas)
	router.POST("/aulas", aulaHandler.InsertarAula)
}

// Generacion de los objetos que se van a usar en la api
func dependencies() {
	//Definicion de variables de interface
	var database repositories.DB
	var aulaRepository repositories.AulaRepositoryInterface
	var aulaService services.AulaInterface

	//Creamos los objetos reales y los pasamos como parametro
	database = repositories.NewMongoDB()
	aulaRepository = repositories.NewAulaRepository(database)
	aulaService = services.NewAulaService(aulaRepository)
	aulaHandler = handlers.NewAulaHandler(aulaService)
}
