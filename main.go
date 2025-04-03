package main

import (
	"Api_usuarios/db"
	"Api_usuarios/model"
	"fmt"
)

func conectarDb() {
	db.ConectarDb()
}

func consultaDePrueba() {
	var usuarios []model.Usuario
	db.DB.Unscoped().Table("usuarios").Find(&usuarios)

	for _, u := range usuarios {
		fmt.Println("---------------------------")
		fmt.Println("Doc: ", u.GetUserDoc())
		fmt.Println("Nombre: ", u.GetUserName())
		fmt.Println("Mail: ", u.GetUserMail())
		fmt.Println("---------------------------")
	}
}
func main() {
	conectarDb()
	consultaDePrueba()
}
