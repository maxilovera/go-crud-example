package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maxilovera/objects/handlers"
)

var (
	aulaHandler *handlers.AulaHandler
	router      *gin.Engine
)

func main() {
	router = gin.Default()
	//Iniciar objetos de handler
	iniciar()

	mapping()

	router.Run(":8080")
}

func mapping() {
	//Listado de rutas
	router.GET("/aulas", aulaHandler.ObtenerAulas)
}

func iniciar() {
	aulaHandler = handlers.NewAulaHandler()
}
