package routes

import (
	"go-http-api/controllers/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {

	userGroup := route.Group("/user")
	{
		userGroup.GET("/listar", user.ListarPersonas)
		userGroup.GET("/listar/:id", user.ObtenerPersona)
		userGroup.POST("/registro", user.CrearPersona)
		userGroup.PUT("/editar/:id", user.ModificarPersona)
		userGroup.DELETE("/eliminar/:id", user.EliminarPersona)
	}

}
