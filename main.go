package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// import estructura "Estructura/Estructura"
// MENU PRINCIPAL
func menuPrincipal() {
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
		//buscar entre los empleados guardados en la lista simple enlazada y si existe iniciar sesion en menu empleado
		/*Se envia usuario y password por parametro al metodo buscar en la lista simple enlazada y si el usuario y el password
		coinciden entonces devuelve true y inicia sesion, de lo contrario devuelve false y muestra el mensaje*/
		//menuEmpleado()
		//si no coincide entonces mostrar el siguiente mensaje
		fmt.Println("El usuario no existe")
	}
}

// MENU ADMINISTRADOR Y SUS FUNCIONES
func menuAdministrador() {
	var opcion int
	var ruta string
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
			fmt.Print("Porfavor ingrese la ruta del archivo")
			fmt.Scanln(&ruta)
			cargarEmpleados(ruta)
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

func cargarEmpleados(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error al abrir el archivo")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	//reader.FieldsPerRecord = 2
	//reader.Comment = '#'
	encabezado := true
	for {
		linea, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No se pudo leer la línea del archivo")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		fmt.Println("Id: ", linea[0], "Nombre: ", linea[1], "Cargo: ", linea[2], "Password: ", linea[3])
	}
}

// MENU EMPLEADO Y SUS FUNCIONES
func menuEmpleado() {
	var opcion int
	for opcion != 3 {
		fmt.Println(`
--------- EDD Creative idEmpleado ---------
1. Ver Imagenes Cargadas
2. Realizar Pedido
3. Cerrar Sesion
-----------------------------------------------------
Elige una opción:`)

		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Print("Porfavor ingrese la ruta del archivo")
		case 2:
			fmt.Print("Estoy en realizar pedido")
		}
	}
}

// METODO MAIN
func main() {
	menuPrincipal()
	/*
		listaSimple := &estructura.Lista_simple{Inicio: nil, Longitud: 0}
		listaSimple.Insertar("1229", "jaquelin Gomez", "Diseño", "1229_Diseño")
		listaSimple.Insertar("3607", "Yadira Ruiz", "Diseño", "3607_Diseño")
		listaSimple.Insertar("3518", "Paula Fuentes", "Ventas", "3518_Ventas")
		listaSimple.Insertar("1211", "karla Alvarez", "Ventas", "1211_Ventas")
		listaSimple.Mostrar()
	*/
}
