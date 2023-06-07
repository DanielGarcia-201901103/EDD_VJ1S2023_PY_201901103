package main

import (
	estructura "Estructura/Estructura"
	"encoding/csv"
	"fmt"
	"os"
)

// variables globales
var listaSimple = estructura.NewListaSimple()
var listaDoble = estructura.NewListaDoble()

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

		validandoExistencia := listaSimple.Validar(usuario, password)

		if validandoExistencia == true {

			menuEmpleado(usuario)
		} else {
			fmt.Println("El usuario no existe")
		}
		//buscar entre los empleados guardados en la lista simple enlazada y si existe iniciar sesion en menu empleado
		/*Se envia usuario y password por parametro al metodo buscar en la lista simple enlazada y si el usuario y el password
		coinciden entonces devuelve true y inicia sesion, de lo contrario devuelve false y muestra el mensaje*/
		//menuEmpleado()
		//si no coincide entonces mostrar el siguiente mensaje
		//fmt.Println("El usuario no existe")
	}
}

// MENU ADMINISTRADOR Y SUS FUNCIONES
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
			cargarEmpleados()
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

func cargarEmpleados() {
	var ruta string
	fmt.Println("Ingrese la ruta del archivo: ")
	fmt.Scanln(&ruta)

	// Abre el archivo CSV
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Crea un nuevo lector CSV
	reader := csv.NewReader(file)

	// Lee todas las líneas del archivo
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	// Itera sobre las líneas y muestra los datos
	for _, line := range lines {
		if line[0] != "id" {
			//fmt.Println(line[0], " ", line[1], " ", line[2], " ", line[3])
			listaSimple.Insertar(line[0], line[1], line[2], line[3])
		}
	}
	listaSimple.Mostrar()
}

// MENU EMPLEADO Y SUS FUNCIONES
func menuEmpleado(usuario string) {
	var opcion int
	for opcion != 3 {
		fmt.Printf(`
--------- EDD Creative %s ---------
1. Ver Imagenes Cargadas
2. Realizar Pedido
3. Cerrar Sesion
-----------------------------------------------------
Elige una opción:`, usuario)

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
	//menuAdministrador()
	/*  ArchivoEmpleados.csv
	listaSimple := &estructura.Lista_simple{Inicio: nil, Longitud: 0}
	listaSimple.Insertar("1229", "jaquelin Gomez", "Diseño", "1229_Diseño")
	listaSimple.Insertar("3607", "Yadira Ruiz", "Diseño", "3607_Diseño")
	listaSimple.Insertar("3518", "Paula Fuentes", "Ventas", "3518_Ventas")
	listaSimple.Insertar("1211", "karla Alvarez", "Ventas", "1211_Ventas")
	listaSimple.Mostrar()
	*/
}
