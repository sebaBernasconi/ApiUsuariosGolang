package main

import (
	"Api_usuarios/controller"
	"Api_usuarios/service"
	"github.com/gin-gonic/gin"
)

// Instanciando el servicio y el controller
var (
	usuarioService    service.UsuarioService       = service.New()
	usuarioController controller.UsuarioController = controller.New(service.UsuarioServiceimpl{})
)

func main() {
	server := gin.Default()

	server.GET("/usuarios", func(c *gin.Context) {
		c.JSON(200, usuarioController.FindAll())
	})

	// El context de gin es una estructura que tiene  codigo de respuesta, request, response, etc
	// por eso cuando le pasas el contexto en el Save(), ya agarra el body de la request
	server.POST("/usuarioNuevo", func(c *gin.Context) {
		c.JSON(200, usuarioController.Save(c))
	})
	server.Run(":8080")
}
