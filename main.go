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

func cargarUsuario(doc int, nombre string, mail string) {
	newUsr := model.Usuario{Documento: doc, Nombre: nombre, Mail: mail}

	db.DB.Omit("created_at", "updated_at", "deleted_at").Create(&newUsr)
	fmt.Println("seee")

}

func main() {
	conectarDb()
	cargarUsuario(11222333, "Rama", "rama@gmail.com")

	consultaDePrueba()
}
