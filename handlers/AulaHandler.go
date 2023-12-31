package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxilovera/go-crud-example/dto"
	"github.com/maxilovera/go-crud-example/services"
	"github.com/maxilovera/go-crud-example/utils"
)

type AulaHandler struct {
	aulaService services.AulaInterface
}

func NewAulaHandler(aulaService services.AulaInterface) *AulaHandler {
	return &AulaHandler{
		aulaService: aulaService,
	}
}

func (handler *AulaHandler) ObtenerAulas(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	//invocamos al método
	filtroNombre := c.Query("nombre")
	aulas := handler.aulaService.ObtenerAulas(filtroNombre)

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:AulaHandler][method:ObtenerAulas][cantidad:%d][user:%s]", len(aulas), user.Codigo)

	c.JSON(http.StatusOK, aulas)
}

func (handler *AulaHandler) ObtenerAulaPorID(c *gin.Context) {
	id := c.Param("id")
	aulas := handler.aulaService.ObtenerAulaPorID(id)

	c.JSON(http.StatusOK, aulas)
}

func (handler *AulaHandler) InsertarAula(c *gin.Context) {
	var aula dto.Aula

	if err := c.ShouldBindJSON(&aula); err != nil {
		// Si hay un error en el JSON, devuelve un error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := handler.aulaService.InsertarAula(&aula)

	c.JSON(http.StatusCreated, resultado)
}

func (handler *AulaHandler) ModificarAula(c *gin.Context) {
	var aula dto.Aula

	if err := c.ShouldBindJSON(&aula); err != nil {
		// Si hay un error en el JSON, devuelve un error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aula.ID = c.Param("id")
	resultado := handler.aulaService.ModificarAula(&aula)

	c.JSON(http.StatusCreated, resultado)
}

func (handler *AulaHandler) EliminarAula(c *gin.Context) {
	id := c.Param("id")
	aulas := handler.aulaService.EliminarAula(id)

	c.JSON(http.StatusOK, aulas)
}
