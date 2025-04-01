package model

import (
	"fmt"
	"time"
)

type Usuario struct {
	documento           int
	nombre              string
	mail                string
	ultimaActualizacion time.Time
}

// Constructor
func NuevoUsuario(doc int, nombre string, mail string, ultimaActualizacion time.Time) *Usuario {
	return &Usuario{documento: doc, nombre: nombre, mail: mail, ultimaActualizacion: ultimaActualizacion}
}

func (u *Usuario) ModificarMail(nueviMail string) {
	u.mail = nueviMail
	u.ultimaActualizacion = time.Now()
	fmt.Println("Mail actualizado con exito")
}

func (u Usuario) MostrarDatos() {
	fmt.Println("Doc: ", u.documento)
	fmt.Println("Nombre: ", u.nombre)
	fmt.Println("Mail: ", u.mail)
	fmt.Println("Ultima Actualizacion: ", u.ultimaActualizacion)

}
