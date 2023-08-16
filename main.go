package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maxilovera/go-crud-example/handlers"
)

func main() {
	r := gin.Default()

	r.GET("/personas", handlers.ObtenerPersonas)
	r.GET("/personas/:id", handlers.ObtenerPersonaPorId)
	r.POST("/personas", handlers.CrearPersona)

	r.Run(":8080")
}
