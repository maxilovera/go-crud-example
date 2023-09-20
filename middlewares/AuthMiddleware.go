package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxilovera/go-crud-example/clients"
	"github.com/maxilovera/go-crud-example/utils"
)

type AuthMiddleware struct {
	authClient clients.AuthClientInterface
}

func NewAuthMiddleware(authClient clients.AuthClientInterface) *AuthMiddleware {
	return &AuthMiddleware{
		authClient: authClient,
	}
}

// Este middleware se ejecuta en el grupo de rutas privadas.
func (auth *AuthMiddleware) ValidateToken(c *gin.Context) {
	//Se obtiene el header necesario con nombre "Authorization"
	authToken := c.GetHeader("Authorization")

	if authToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token no encontrado"})
		return
	}

	//Obtener la informacion del usuario a partir del token desde el servicio externo
	user, err := auth.authClient.GetUserInfo(authToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autorizado"})
		return
	}

	//Seteamos los datos del usuario logueado en el contexto de GIN.
	utils.SetUserInContext(c, user)

	c.Next()
}
