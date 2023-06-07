package main

import (
	estructura "Estructura/Estructura"
	"encoding/csv"
	"fmt"
	"os"
)

//import estructura "Estructura/Estructura"

func login() {
	var opcion int
	for opcion != 2 {
		fmt.Println(`
--------- Login ---------
1. Iniciar Sesion
2. Salir del Sistema
-------------------------
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
	for opcion != 6 {
		fmt.Println(`
--------- Dashboard Administrador 201901103 ---------
1. Cargar Empleados
2. Cargar Imagenes
3. Cargar Usuarios
4. Actualizar Cola
5. Reportes Estructuras
6. Cerrar Sesion
-----------------------------------------------------
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
		}
	}
}

func cargarArchivo() {
	file, err := os.Open("archivosPrueba/imagenes.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	//reader.Comma = ','
	reader.FieldsPerRecord = 2
	reader.Comment = '#'

	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}
}

func main() {
	//login()

	listaSimple := &estructura.Lista_simple{Inicio: nil, Longitud: 0}
	listaSimple.Insertar("1229", "jaquelin Gomez", "Diseño", "1229_Diseño")
	listaSimple.Insertar("3607", "Yadira Ruiz", "Diseño", "3607_Diseño")
	listaSimple.Insertar("3518", "Paula Fuentes", "Ventas", "3518_Ventas")
	listaSimple.Insertar("1211", "karla Alvarez", "Ventas", "1211_Ventas")
	listaSimple.Mostrar()

}
