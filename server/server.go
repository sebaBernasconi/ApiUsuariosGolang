package main

import (
	"Api_usuarios/controller"
	"Api_usuarios/service"
	"github.com/gin-gonic/gin"
)

var (
	usuarioService    service.UsuarioService       = service.New()
	usuarioController controller.UsuarioController = controller.New(service.UsuarioServiceimpl{})
)

func main() {
	server := gin.Default()

	server.GET("/usuarios", func(c *gin.Context) {
		c.JSON(200, usuarioController.FindAll())
	})

	server.POST("/usuarioNuevo", func(c *gin.Context) {
		c.JSON(200, usuarioController.Save(c))
	})
	server.Run(":8080")
}
