package controller

import (
	"Api_usuarios/model"
	"Api_usuarios/service"
	"github.com/gin-gonic/gin"
)

// Interfaz de los metodos que tiene que implementar el controller y lo que retornan
type UsuarioController interface {
	FindAll() []model.Usuario
	Save(c *gin.Context) model.Usuario
}

// Instanciando el servicio (interfaz)
type controller struct {
	service service.UsuarioService
}

// Constructor del controller que recibe la estructura del servicio y devuelve
// un puntero a la estructura del controller
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
	controller.service.Save(usuario)
	return usuario
}
