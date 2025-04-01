package main

import (
	"fmt"
	"time"

	"Api_usuarios/model"
)

func main() {

	u := model.NuevoUsuario(24567864, "Sebas", "sebas@gmail.com", time.Now())

	u.MostrarDatos()
	fmt.Println("------------------------------------------------------")

	u.ModificarMail("nuevo@gmail.com")
	u.MostrarDatos()
}
