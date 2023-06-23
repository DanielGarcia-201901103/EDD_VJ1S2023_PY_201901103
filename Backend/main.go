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

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

type RespImagen struct {
	Imagenbase64 string
	Nombre       string
}

// variables globales
var listaSimple = estructura.NewListaSimple()
var listaDoble = estructura.NewListaDoble()
var listaCircular = estructura.NewListaCircular()
var clientesCola = estructura.NewCola()
var pedidosPila = estructura.NewPila()
var arbol *estructura.ArbolAVL

//var matrizImages = &estructura.Matriz{Raiz: &estructura.NodoMatriz{PosicionX: -1, PosicionY: -1, Color: "RAIZ"}}

// import estructura "Estructura/Estructura"
// MENU PRINCIPAL
/*
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
*/
func sesion(usuario string, password string) string {
	if usuario == "ADMIN_201901103" && password == "Admin" {
		//menuAdministrador()
		return "Administrador 201901103"
	} else {

		validandoExistencia := listaSimple.Validar(usuario, password)

		if validandoExistencia == true {
			menuEmpleado(usuario)
			return usuario
		}
	}
	return "No"
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
	//utf8Reader := transform.NewReader(file, unicode.UTF8.NewDecoder())

	// Crea un nuevo lector CSV
	reader := csv.NewReader(file)
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

type DatosUser struct {
	Usuario  string `json: "Usuario"`
	Password string `json: "Password"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		//return c.SendString("Hello, World!")
		return c.JSON(&fiber.Map{
			"data": "hola",
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		jsonData := new(DatosUser)
		if err := c.BodyParser(jsonData); err != nil {
			return err
		}
		usuarioRecibido := jsonData.Usuario
		passwordRecibido := jsonData.Password
		fmt.Println(usuarioRecibido)
		fmt.Println(passwordRecibido)
		validacionIniciar := sesion(usuarioRecibido, passwordRecibido)

		if validacionIniciar == "Administrador 201901103" {
			fmt.Print("Administrador 201901103")
			return c.JSON(&fiber.Map{
				"estado": "Administrador 201901103",
			})
		}
		if validacionIniciar != "No" {
			fmt.Print("Cualquier usuario")
			return c.JSON(&fiber.Map{
				"estado": "SI",
			})
		} else {
			fmt.Print("Usuario o contraseña incorrectos")
			return c.JSON(&fiber.Map{
				"estado": "NO",
			})
		}
	})

	app.Listen(":5000")
	//https://github.com/gofiber/fiber
}

// METODO MAIN
/*
func main() {
	//menuPrincipal()
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/agregar-arbol", func(c *fiber.Ctx) error {
		var arbol estructura.NodoAVL
		c.BodyParser(&arbol)
		fmt.Println(arbol.Data)
		return c.Json(&fiber.Map{
			"data": "hola",
		})
	})
	app.Listen(":3003")

		https://github.com/gofiber/fiber
			arbol = &estructura.ArbolAVL{Raiz: nil}
			r := mux.NewRouter()
			r.HandleFunc("/", MostrarArbol).Methods("GET")
			r.HandleFunc("/agregar-arbol", AgregarArbol).Methods("POST")
			r.HandleFunc("/reporte-arbol", MandarReporte).Methods("GET")
			log.Fatal(http.ListenAndServe(":3001", r))

}
*/
/*
func MostrarArbol(w http.ResponseWriter, req *http.Request) {
	//Esto nos verifica que le estamos enviando al servidor una respuesta de tipo JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&arbol)
}

func AgregarArbol(w http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	var nuevoNodo estructura.NodoAVL
	if err != nil {
		fmt.Fprintf(w, "No valido")
	}
	json.Unmarshal(reqBody, &nuevoNodo)
	fmt.Println(nuevoNodo.Data)
	arbol.InsertarElemento(nuevoNodo.Data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoNodo)
}

func MandarReporte(w http.ResponseWriter, req *http.Request) {
	arbol.Graficar()
	var imagen RespImagen = RespImagen{Nombre: "arbolAVL.jpg"}
	//INICIO
	imageBytes, err := ioutil.ReadFile(imagen.Nombre)
	if err != nil {
		fmt.Fprintf(w, "Imagen No Valida")
		return
	}
	// Codifica los bytes de la imagen en base64
	imagen.Imagenbase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)

	//data:image/jpg;base64,ABC
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(imagen)
}*/

/*
https://drive.google.com/file/d/1Mu40-ZEfP-CMmgPoNtdIBWNoCNng1JYb/view
https://github.com/CristianMejia2198/EDD_1S_JUNIO_2023
https://www.markdownguide.org/basic-syntax/#images-1
https://markdown.es/sintaxis-markdown/
https://github.com/CristianMejia2198/S1EDD-C/tree/main/Clase6


FASE 2
sistema principal
agregar opcion
	historial facturas

biller: id empleado
*/
