package main

import (
	estructura "Estructura/Estructura"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type RespImagen struct {
	Imagenbase64 string
	Nombre       string
}

type DatosUser struct {
	Usuario  string `json: "Usuario"`
	Password string `json: "Password"`
}

type URLempleado struct {
	Ruta string `json: "Ruta"`
}

type Pedi struct {
	ID     int    `json:"id_cliente"`
	Imagen string `json:"imagen"`
}

type Filt struct {
	Tipo   string `json:"Tipo"`
	Imagen string `json:"Imagen"`
}

// variables globales
var listaSimple = estructura.NewListaSimple()
var listaDoble = estructura.NewListaDoble()
var listaCircular = estructura.NewListaCircular()
var clientesCola = estructura.NewCola()
var pedidosPila = estructura.NewPila()
var arbol estructura.ArbolAVL

func sesion(usuario string, password string) string {
	if usuario == "ADMIN_201901103" && password == "Admin" {
		//menuAdministrador()
		return "Administrador 201901103"
	} else {
		validandoExistencia := listaSimple.Validar(usuario, password)
		if validandoExistencia {
			//menuEmpleado(usuario)
			return usuario
		}
	}
	return "No"
}

// MENU ADMINISTRADOR Y SUS FUNCIONES

func cargarEmpleados(ruta string) bool {
	// Abre el archivo CSV
	file, err := os.Open(ruta)
	if err != nil {
		//fmt.Println("Error al abrir el archivo:", err)
		return false
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
		//fmt.Println("Error al leer el archivo:", err)
		return false
	}

	// Itera sobre las líneas y muestra los datos
	for _, line := range lines {
		if line[0] != "id" {
			//fmt.Println(line[0], " ", line[1], " ", line[2], " ", line[3])
			listaSimple.Insertar(strings.TrimSpace(line[0]), strings.TrimSpace(line[1]), strings.TrimSpace(line[2]), strings.TrimSpace(line[3]))
		}
	}
	//listaSimple.Mostrar()
	//fmt.Println("Carga exitosa")
	return true
}

func previaVisualizacion(nameImagen string, tipoFiltro string) {
	var matrizImages = &estructura.Matriz{Raiz: &estructura.NodoMatriz{PosicionX: -1, PosicionY: -1, Color: "RAIZ"}}
	matrizImages.LeerInicial("csv/"+nameImagen+"/inicial.csv", nameImagen)
	matrizImages.GenerarImagen(nameImagen, "Original")

	if tipoFiltro == "escalaGris" {
		matrizImages.FiltroEscalaGris(nameImagen)
	} else if tipoFiltro == "escalaNegativo" {
		matrizImages.FiltroNegativo(nameImagen)
	} else if tipoFiltro == "espejoX" {
		matrizImages.EspejoX()
		matrizImages.GenerarImagen(nameImagen, tipoFiltro)
	} else if tipoFiltro == "espejoY" {
		matrizImages.EspejoY()
		matrizImages.GenerarImagen(nameImagen, tipoFiltro)
	} else if tipoFiltro == "dobleEspejo" {
		matrizImages.EspejoDoble()
		matrizImages.GenerarImagen(nameImagen, tipoFiltro)
	} else {
		fmt.Println("El espejo no existe")
	}

	matrizImages = &estructura.Matriz{Raiz: nil}
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

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		//return c.SendString("Hello, World!")
		return c.JSON(&fiber.Map{
			"data": "Bienvenido a EDD Creative",
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		jsonData := new(DatosUser)
		if err := c.BodyParser(jsonData); err != nil {
			return err
		}
		usuarioRecibido := jsonData.Usuario
		passwordRecibido := jsonData.Password
		//fmt.Println(usuarioRecibido)
		//fmt.Println(passwordRecibido)
		validacionIniciar := sesion(usuarioRecibido, passwordRecibido)

		if validacionIniciar == "Administrador 201901103" {
			//fmt.Print("Administrador 201901103")
			return c.JSON(&fiber.Map{
				"data": "Administrador",
			})
		}
		if validacionIniciar != "No" {
			//fmt.Print("Cualquier usuario")
			return c.JSON(&fiber.Map{
				"data": "SI",
			})
		} else {
			//fmt.Print("Usuario o contraseña incorrectos")
			return c.JSON(&fiber.Map{
				"data": "NO",
			})
		}
	})

	app.Post("/cargaEmpleados", func(c *fiber.Ctx) error {
		//return c.SendString("Hello, World!")
		jsonUrl := new(URLempleado)
		if err := c.BodyParser(jsonUrl); err != nil {
			return err
		}
		rutaRecibida := jsonUrl.Ruta
		validacionleer := cargarEmpleados(rutaRecibida)

		if validacionleer {
			return c.JSON(&fiber.Map{
				"data": "archivo cargado correctamente",
			})
		}

		return c.JSON(&fiber.Map{
			"data": "error al cargar archivo",
		})
	})

	app.Post("/cargarPedidos", func(c *fiber.Ctx) error {
		jsonUrl := new(URLempleado)
		if err := c.BodyParser(jsonUrl); err != nil {
			return err
		}
		rutaRecibida := jsonUrl.Ruta
		validacionleer := cargarJson(rutaRecibida)
		arbol.InOrder(clientesCola)
		if validacionleer {
			return c.JSON(&fiber.Map{
				"data": "archivo cargado correctamente",
			})
		}

		return c.JSON(&fiber.Map{
			"data": "error al cargar archivo",
		})
	})

	app.Get("/Reportes", func(c *fiber.Ctx) error {
		//return c.SendString("Hello, World!")
		arbol.Graficar()
		clientesCola.ReporteCola()
		var imagen RespImagen = RespImagen{Nombre: "arbolAVL.jpg"}
		//INICIO
		imageBytes, err := ioutil.ReadFile(imagen.Nombre)
		fmt.Println(imagen.Nombre)
		if err != nil {
			//fmt.Fprintf(w, "Imagen No Valida")
			return c.JSON(&fiber.Map{
				"data": "error en imagen",
			})
		}
		// Codifica los bytes de la imagen en base64
		imagen.Imagenbase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)

		return c.JSON(&fiber.Map{
			"data": imagen.Imagenbase64,
		})
	})

	app.Get("/clienteObtener", func(c *fiber.Ctx) error {
		idcolaClientes := clientesCola.ObtenerClienteId()
		nameColaClientes := clientesCola.ObtenerClienteName()
		longi := clientesCola.ObtenerLongitud()
		if longi != 0 {
			return c.JSON(&fiber.Map{
				"data":   idcolaClientes,
				"imagen": nameColaClientes,
			})
		}
		return c.JSON(&fiber.Map{
			"data": "sin clientes por atender",
		})

	})

	app.Get("/clienteBorrar", func(c *fiber.Ctx) error {
		clientesCola.Descolar()
		return c.JSON(&fiber.Map{
			"data": "cliente atendido",
		})
	})

	//PROBANDO LOS FILTROS
	app.Post("/filtroNegativo", func(c *fiber.Ctx) error {
		//return c.SendString("Hello, World!")
		img := new(Filt)
		if err := c.BodyParser(img); err != nil {
			return err
		}
		imgRecibida := img.Imagen
		tipoRecibido := img.Tipo

		previaVisualizacion(imgRecibida, tipoRecibido)

		return c.JSON(&fiber.Map{
			"data": "archivo cargado correctamente",
		})
	})

	app.Listen(":5000")
	//https://github.com/gofiber/fiber
}

/*
pedidos.json

csv\empleados.csv


4269,Paula Fuentes,Ventas,2576_Ventas
4364,Maria Tux,Ventas,4364_Ventas



algoritmo para filtros

elegir una imagen

1. generar imagen original de la misma
2. preservar la imagen original
3. aplicar el filtro a la imagen
4. aplicar el segundo filtro a la imagen a
   la que ya se le habia aplicado filtro
*/

func cargarJson(ruta string) bool {
	file, err := os.Open(ruta)
	if err != nil {
		return false
	}
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return false
	}

	var objeto struct {
		Pedidos []Pedi `json:"pedidos"`
	}

	err = json.Unmarshal(byteValue, &objeto)
	if err != nil {
		return false
	}

	for _, pedi := range objeto.Pedidos {
		idTempo := pedi.ID
		imagenTempo := pedi.Imagen
		arbol.InsertarElemento(idTempo, imagenTempo)
	}
	return true
}

/*
https://drive.google.com/file/d/1Mu40-ZEfP-CMmgPoNtdIBWNoCNng1JYb/view
https://github.com/CristianMejia2198/EDD_1S_JUNIO_2023
https://www.markdownguide.org/basic-syntax/#images-1
https://markdown.es/sintaxis-markdown/
https://github.com/CristianMejia2198/S1EDD-C/tree/main/Clase6
https://w3.unpocodetodo.info/canvas/blancoynegro.php

https://github.com/gofiber/fiber
FASE 2
sistema principal
agregar opcion
	historial facturas

biller: id empleado
*/
