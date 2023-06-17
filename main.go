package main

import (
	estructura "Estructura/Estructura"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// variables globales
var listaSimple = estructura.NewListaSimple()
var listaDoble = estructura.NewListaDoble()
var listaCircular = estructura.NewListaCircular()
var clientesCola = estructura.NewCola()
var pedidosPila = estructura.NewPila()

//var matrizImages = &estructura.Matriz{Raiz: &estructura.NodoMatriz{PosicionX: -1, PosicionY: -1, Color: "RAIZ"}}

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
Seleccione una opción:`)

		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			sesion()
		}
	}
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
			fmt.Println("El usuario no existe o ingresó mal el usuario.")
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
Seleccione una opción:`)

		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			cargarEmpleados()
		case 2:
			cargarImagenes()
		case 3:
			cargarClientes()
		case 4:
			cargarActualizarCola()
		case 5:
			listaSimple.ReporteSimple()
			listaDoble.ReporteDoble()
			listaCircular.ReporteCircular()
			clientesCola.ReporteCola()
			//fmt.Print("Estoy en reportes estructuras")
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

	// Crea un lector con transformador UTF-8
	utf8Reader := transform.NewReader(file, unicode.UTF8.NewDecoder())

	// Crea un nuevo lector CSV
	reader := csv.NewReader(utf8Reader)
	reader.Comma = ','

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
			listaSimple.Insertar(strings.TrimSpace(line[0]), strings.TrimSpace(line[1]), strings.TrimSpace(line[2]), strings.TrimSpace(line[3]))
		}
	}
	//listaSimple.Mostrar()
}

func cargarImagenes() {
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

	// Crea un lector con transformador UTF-8
	utf8Reader := transform.NewReader(file, unicode.UTF8.NewDecoder())

	// Crea un nuevo lector CSV
	reader := csv.NewReader(utf8Reader)
	reader.Comma = ','
	encabezado := true

	for {
		lines, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error al leer la linea del archivo")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		listaDoble.Insertar(strings.TrimSpace(lines[0]), strings.TrimSpace(lines[1]))
	}
	//listaDoble.MostrarAscendente()
	//listaDoble.MostrarDescendente()
}

func cargarClientes() {
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

	// Crea un lector con transformador UTF-8
	utf8Reader := transform.NewReader(file, unicode.UTF8.NewDecoder())

	// Crea un nuevo lector CSV
	reader := csv.NewReader(utf8Reader)
	reader.Comma = ','
	encabezado := true

	for {
		lines, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error al leer la linea del archivo")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		listaCircular.Insertar(strings.TrimSpace(lines[0]), strings.TrimSpace(lines[1]))
	}
	//listaCircular.Mostrar()
}

func cargarActualizarCola() {
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

	// Crea un lector con transformador UTF-8
	utf8Reader := transform.NewReader(file, unicode.UTF8.NewDecoder())

	// Crea un nuevo lector CSV
	reader := csv.NewReader(utf8Reader)
	reader.Comma = ','
	encabezado := true

	for {
		lines, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error al leer la linea del archivo")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		clientesCola.Encolar(strings.TrimSpace(lines[0]), strings.TrimSpace(lines[1]))
	}

}

// MENU EMPLEADO Y SUS FUNCIONES
func menuEmpleado(usuario string) {
	var opcion int
	for opcion != 4 {
		fmt.Printf(`
--------- EDD Creative %s ---------
1. Ver Imagenes Cargadas
2. Realizar Pedido
3. Capas
4. Cerrar Sesion
-----------------------------------------------------
Seleccione una opción:`, usuario)

		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			nameImagen := visualizarImagenes()
			fmt.Println("La imagen elegida fue: ", nameImagen, "\nMostrando visualizacion previa")
			previaVisualizacion(nameImagen)
		case 2:
			realizarPedidos(usuario)
			pedidosPila.ReportePila()
			pedidosPila.ReporteJson()
		case 3:
			nameImagen := visualizarImagenes()
			fmt.Println("La imagen elegida fue: ", nameImagen, "\nMostrando visualizacion previa")
			realizarCapa(nameImagen)
		}

	}
}

func visualizarImagenes() string {
	var opcion int
	fmt.Println("\n###################Listado de Imagenes###################")
	listaDoble.ListarDatos()
	fmt.Println("\n Seleccione una opción:")
	fmt.Scanln(&opcion)
	nameImagen := listaDoble.BuscarImagen(strconv.Itoa(opcion))
	return nameImagen
	//Falta la opcion de visualizar la imagen
}

func previaVisualizacion(nameImagen string) {
	var matrizImages = &estructura.Matriz{Raiz: &estructura.NodoMatriz{PosicionX: -1, PosicionY: -1, Color: "RAIZ"}}
	matrizImages.LeerInicial("csv/"+nameImagen+"/inicial.csv", nameImagen)
	matrizImages.GenerarImagen(nameImagen)
	matrizImages = &estructura.Matriz{Raiz: nil}
}

func realizarPedidos(usuario string) {
	for {
		idcolaClientes := clientesCola.ObtenerClienteId()
		nameColaClientes := clientesCola.ObtenerClienteName()
		longi := clientesCola.ObtenerLongitud()
		if longi != 0 {
			fmt.Println("\nAtendiendo al cliente con id: ", idcolaClientes, " y nombre: ", nameColaClientes)

			if strings.ToUpper(idcolaClientes) == "X" {
				// CUANDO ES IGUAL A X VALIDAR UN ID RANDOM Y
				for {
					valor := (rand.Intn(10000)) + 1000

					existe := listaCircular.ValidarRepetidos(strconv.Itoa(valor))
					if existe == true {
						//repetir el aleatorio y no guardar nada
					} else {
						// guardar el aleatorio como nuevo id y agregarlo a la lista circular junto al nombre del cliente
						nombreImagenElegida := visualizarImagenes()
						//Sino existe en la lista circular agregar al cliente en la lista circular
						listaCircular.Insertar(strconv.Itoa(valor), nameColaClientes)
						pedidosPila.Push(strconv.Itoa(valor), usuario, nombreImagenElegida)
						//agregar el id del cliente, id del empleado, y nombre de la imagen elegida
						fmt.Println("\nEl nuevo id: ", strconv.Itoa(valor), "corresponde al cliente: ", nameColaClientes)
						clientesCola.Descolar()
						break
					}
				}

			} else {
				existe := listaCircular.ValidarRepetidos(strings.TrimSpace(idcolaClientes))
				if existe == true {
					// si el cliente existe en la lista circular de clientes
					nombreImagenElegida := visualizarImagenes()
					pedidosPila.Push(idcolaClientes, usuario, nombreImagenElegida)
					//agregar el id del cliente, id del empleado, y nombre de la imagen elegida
					clientesCola.Descolar()
				} else {
					nombreImagenElegida := visualizarImagenes()
					//Sino existe en la lista circular agregar al cliente en la lista circular
					listaCircular.Insertar(idcolaClientes, nameColaClientes)
					pedidosPila.Push(idcolaClientes, usuario, nombreImagenElegida)
					//agregar el id del cliente, id del empleado, y nombre de la imagen elegida
					clientesCola.Descolar()
				}

			}
			fmt.Println("\nFinaliza atención a cliente actual y quedan:", strconv.Itoa(longi-1))
		} else {
			break
		}
	}
	/*
		1.Buscar el primer cliente y obtener su id en la cola
		y retornarlo

		2.validar si el id es igual a x, entonces
		crear un id aleatorio y que sea diferente a los id existentes en la lista circular y luego retornar el id
		y agregar el nombre y el nuevo id a la lista circular
		y imprimir en pantalla el id del cliente y su respectivo nombre
		3. si no es igual a x, entonces continuar con el paso 4
		4. Tomar el Id de cliente, id de empleado, nombre de la imagen y guardarlo dentro de la pila

	*/
}
func realizarCapa(nameImagen string) {
	var matrizImages1 = &estructura.Matriz{Raiz: &estructura.NodoMatriz{PosicionX: -1, PosicionY: -1, Color: "RAIZ"}}
	var listaCapasMatriz = estructura.NewListaSimpleCapa()
	matrizImages1.LeerInicial1("csv/"+nameImagen+"/inicial.csv", nameImagen, listaCapasMatriz)
	matrizImages1 = &estructura.Matriz{Raiz: nil}

	var opcion int
	fmt.Println("\n=================Listado de Capas=================")
	listaCapasMatriz.ListarDatosCapa()
	fmt.Println("\n Seleccione una opción:")
	fmt.Scanln(&opcion)
	nameCapa := listaCapasMatriz.BuscarCapa(strconv.Itoa(opcion))
	matrizImages1 = &estructura.Matriz{Raiz: &estructura.NodoMatriz{PosicionX: -1, PosicionY: -1, Color: "RAIZ"}}
	matrizImages1.LeerInicialYCapaElegida("csv/"+nameImagen+"/inicial.csv", nameImagen, nameCapa)
	matrizImages1 = &estructura.Matriz{Raiz: nil}
	/*
		leer el archivo inicial con cada capa
		leer la capa de configuración
		guardar cada una de las capas en una lista
		recorrer la lista y mostrar las opciones de las capas, luego preguntar cual capa elige
		enviar el nombre de la capa al metodo LeerArchivo
	*/
}

// METODO MAIN
func main() {
	menuPrincipal()
}

/*
https://drive.google.com/file/d/1Mu40-ZEfP-CMmgPoNtdIBWNoCNng1JYb/view
https://github.com/CristianMejia2198/EDD_1S_JUNIO_2023
https://www.markdownguide.org/basic-syntax/#images-1
https://markdown.es/sintaxis-markdown/
https://github.com/CristianMejia2198/S1EDD-C/tree/main/Clase6
*/
