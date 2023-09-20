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
	//cliente para api externa
	var authClient clients.AuthClientInterface
	authClient = clients.NewAuthClient()
	//creacion de middleware de autenticacion
	authMiddleware := middlewares.NewAuthMiddleware(authClient)

	//Listado de rutas
	group := router.Group("/aulas")
	//Uso del middleware para todas las rutas del grupo
	group.Use(authMiddleware.ValidateToken)

	group.GET("/", aulaHandler.ObtenerAulas)
	group.GET("/:id", aulaHandler.ObtenerAulaPorID)
	group.POST("/", aulaHandler.InsertarAula)
	group.PUT("/:id", aulaHandler.ModificarAula)
	group.DELETE("/:id", aulaHandler.EliminarAula)

	groupE := router.Group("/escuelas")
	groupE.GET("/", aulaHandler.ObtenerAulas)
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
