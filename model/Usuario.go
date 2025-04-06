package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model

	Documento int
	Nombre    string
	Mail      string
}

// Constructor
func NuevoUsuario(doc int, Nombre string, Mail string) *Usuario {
	return &Usuario{Documento: doc, Nombre: Nombre, Mail: Mail}
}

func (u *Usuario) ModificarMail(nueviMail string) {
	u.Mail = nueviMail
	fmt.Println("Mail actualizado con exito")
}

func (u Usuario) MostrarDatos() {
	fmt.Println("Doc: ", u.Documento)
	fmt.Println("Nombre: ", u.Nombre)
	fmt.Println("Mail: ", u.Mail)

}
