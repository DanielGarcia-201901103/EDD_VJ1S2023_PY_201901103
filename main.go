package main

import "fmt"

//import estructura "Estructura/Estructura"

func login() {
	var opcion int
	for opcion != 2 {
		fmt.Println(`
--------- Login ---------
	1. Iniciar Sesion
	2. Salir del Sistema
Elige una opción:`)

		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			sesion()
		case 2:
			break
		}
	}

	//fmt.Scanf("%s\n", &name)
	//fmt.Printf("Bienvenido %s", name)
}

func sesion() {
	var usuario string
	var password string
	fmt.Println("\nIngrese Usuario: ")
	fmt.Scanln(&usuario)
	fmt.Println("Password: ")
	fmt.Scanln(&password)

	if usuario == "ADMIN_201901103" && password == "Admin" {
		fmt.Println("Bienvenido a admin")
		menuAdministrador()
	} else {
		fmt.Println("El usuario no existe")
	}
}

func menuAdministrador() {
	var opcion int
	for opcion != 7 {
		fmt.Println(`
--------- Dashboard Administrador 201901103 ---------
	1. Cargar Empleados
	2. Cargar Imagenes
	3. Cargar Usuarios
	4. Actualizar Cola
	5. Reportes Estructuras
	6. Regresar
Elige una opción:`)

		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Print("Estoy en cargar empleados")
		case 2:
			fmt.Print("Estoy en cargar imagenes")
		case 3:
			fmt.Print("Estoy en cargar usuarios")
		case 4:
			fmt.Print("Estoy en actualizar cola")
		case 5:
			fmt.Print("Estoy en reportes estructuras")
		case 6:
			login()
		}
	}
}

func main() {
	login()
	/*
		listaSimple := &estructura.Lista{Inicio: nil, Longitud: 0}
		listaSimple.Insertar(1)
		listaSimple.Insertar(3)
		listaSimple.Insertar(5)
		listaSimple.Insertar(7)
		listaSimple.Mostrar()
	*/
}

//video de alura ver
//https://www.youtube.com/watch?v=QUylEZSrRok
