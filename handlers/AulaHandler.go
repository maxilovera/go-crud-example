package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxilovera/go-crud-example/services"
)

type AulaHandler struct {
	aulaService services.AulaService
}

func NewAulaHandler() *AulaHandler {
	return &AulaHandler{
		aulaService: *services.NewAulaService(),
	}
}

func (handler AulaHandler) ObtenerAulas(c *gin.Context) {
	aulas := handler.aulaService.ObtenerAulas()

	contentType := c.Request.Header.Get("content-type")

	if contentType == "application/xml" {
		c.XML(http.StatusOK, aulas)
		return
	}

	c.JSON(http.StatusOK, aulas)
}
