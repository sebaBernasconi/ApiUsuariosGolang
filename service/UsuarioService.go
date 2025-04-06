package service

import "Api_usuarios/model"

// Interfaz con los metodos que voy a tener que implemetntar
type UsuarioService interface {
	Save(usuario model.Usuario) model.Usuario
	FindAll() []model.Usuario
}

// estructura donde voy a guardar los usuarios que cargue que es un slice de estructuras usr
type UsuarioServiceimpl struct {
	usuarios []model.Usuario
}

// Constructor con puntero a la estructura del servicio
func New() UsuarioService {
	return &UsuarioServiceimpl{}
}

func (service *UsuarioServiceimpl) Save(usuario model.Usuario) model.Usuario {
	service.usuarios = append(service.usuarios, usuario)
	return usuario
}

func (service *UsuarioServiceimpl) FindAll() []model.Usuario {
	return service.usuarios

}
