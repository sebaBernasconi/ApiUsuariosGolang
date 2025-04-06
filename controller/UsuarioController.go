package controller

import (
	"Api_usuarios/model"
	"Api_usuarios/service"
	"github.com/gin-gonic/gin"
)

type UsuarioController interface {
	FindAll() []model.Usuario
	Save(c *gin.Context) model.Usuario
}

type controller struct {
	service service.UsuarioService
}

func New(service service.UsuarioServiceimpl) UsuarioController {
	return &controller{
		service: &service,
	}
}

func (controller *controller) FindAll() []model.Usuario {
	return controller.service.FindAll()
}

func (controller *controller) Save(c *gin.Context) model.Usuario {
	var usuario model.Usuario
	c.BindJSON(&usuario)
	controller.servicie.Save(usuario)
	return usuario
}
