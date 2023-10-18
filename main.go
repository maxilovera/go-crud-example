package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/maxilovera/go-crud-example/clients"
	"github.com/maxilovera/go-crud-example/handlers"
	"github.com/maxilovera/go-crud-example/middlewares"
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

	log.Println("Iniciando el servidor...")
	router.Run(":8080")
}

func mappingRoutes() {
	//middleware para permitir peticiones del mismo server localhost
	router.Use(middlewares.CORSMiddleware())

	//cliente para api externa
	var authClient clients.AuthClientInterface
	authClient = clients.NewAuthClient()
	//creacion de middleware de autenticacion
	authMiddleware := middlewares.NewAuthMiddleware(authClient)

	//Listado de rutas
	//Uso del middleware para todas las rutas del grupo
	router.Use(authMiddleware.ValidateToken)

	router.GET("/aulas", aulaHandler.ObtenerAulas)
	router.GET("/aulas/:id", aulaHandler.ObtenerAulaPorID)
	router.POST("/aulas", aulaHandler.InsertarAula)
	router.PUT("/aulas/:id", aulaHandler.ModificarAula)
	router.DELETE("/aulas/:id", aulaHandler.EliminarAula)
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
