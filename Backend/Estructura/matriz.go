package estructura

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type NodoMatriz struct {
	Siguiente *NodoMatriz
	Anterior  *NodoMatriz
	Abajo     *NodoMatriz
	Arriba    *NodoMatriz
	PosicionX int
	PosicionY int
	Color     string
}

type Matriz struct {
	Raiz        *NodoMatriz
	ImageWidth  int
	ImageHeight int
	PixelWidth  int
	PixelHeight int
}

func (m *Matriz) buscarColumna(x int) *NodoMatriz {
	//accede a la raiz
	aux := m.Raiz
	//valida si es diferente de nulo
	for aux != nil {
		// valida si la posicion en x es igual a la posicion de la columna
		if aux.PosicionX == x {
			//retorna el nodo de la matriz
			return aux
		}
		// pasa a la siguiente columna
		aux = aux.Siguiente
	}
	//si es nulo retorna nil
	return nil
}

func (m *Matriz) buscarFila(y int) *NodoMatriz {
	//accede a la raiz de la matriz
	aux := m.Raiz
	//valida si el nodo es diferente de nulo realiza el bucle
	for aux != nil {
		//valida si la posicion de la fila coincide con la posicion y
		if aux.PosicionY == y {
			//retorna el nodo de la matriz
			return aux
		}
		//pasa a la siguiente fila
		aux = aux.Abajo
	}
	//si es nulo retorna nil
	return nil
}

func (m *Matriz) insertarColumna(nuevoNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz {
	// accede al nodo raiz de la matriz
	temp := nodoRaiz
	// se inicializa el pivote que servira
	piv := false
	for { // while(true) [2][2][2][5][5] -> [N]
		//valida si la posicion en x es igual a la nueva posicion en x del nuevo nodo
		if temp.PosicionX == nuevoNodo.PosicionX {
			//la posicion en la fila obtiene la posicion de y
			temp.PosicionY = nuevoNodo.PosicionY
			// al nodo se le asigna el color
			temp.Color = nuevoNodo.Color
			//retorna el nodo
			return temp
		} else if temp.PosicionX > nuevoNodo.PosicionX { //Si la posicion en x es mayor que la posicion del nuevo nodo
			piv = true
			break
		}
		if temp.Siguiente != nil { // si el siguiente es diferente de nulo
			temp = temp.Siguiente // pasa al siguiente nodo
		} else { //cuando el siguiente es nulo
			break //finaliza el bucle
		}
	}
	if piv {
		/*Asumir que nuevo = C1*/
		nuevoNodo.Siguiente = temp          // C2
		temp.Anterior.Siguiente = nuevoNodo // siguiente de raiz ahora es C1
		nuevoNodo.Anterior = temp.Anterior  // Anterior Raiz
		temp.Anterior = nuevoNodo           //
	} else {
		temp.Siguiente = nuevoNodo //nodo siguiente apunta al nuevo nodo
		nuevoNodo.Anterior = temp  // el anterior del nuevo nodo apunta al nodo
	}
	return nuevoNodo //retorna el nuevo nodo
}

func (m *Matriz) insertarFila(nuevoNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz {
	temp := nodoRaiz
	piv := false
	for { //
		if temp.PosicionY == nuevoNodo.PosicionY {
			temp.PosicionX = nuevoNodo.PosicionX
			temp.Color = nuevoNodo.Color
			return temp
		} else if temp.PosicionY > nuevoNodo.PosicionY {
			piv = true
			break
		}
		if temp.Abajo != nil {
			temp = temp.Abajo
		} else {
			break
		}
	}
	if piv {
		/*Asumir que nuevo = C1*/
		nuevoNodo.Abajo = temp         // C2
		temp.Arriba.Abajo = nuevoNodo  // siguiente de raiz ahora es C1
		nuevoNodo.Arriba = temp.Arriba // Anterior Raiz
		temp.Arriba = nuevoNodo        //
	} else {
		temp.Abajo = nuevoNodo
		nuevoNodo.Arriba = temp
	}
	return nuevoNodo
}

func (m *Matriz) nuevaColumna(x int) *NodoMatriz {
	col := "C" + strconv.Itoa(x)                                      // C1
	nuevoNodo := &NodoMatriz{PosicionX: x, PosicionY: -1, Color: col} //encabezado de las columnas
	columna := m.insertarColumna(nuevoNodo, m.Raiz)                   //agrega la columna
	return columna
}

func (m *Matriz) nuevaFila(y int) *NodoMatriz {
	col := "F" + strconv.Itoa(y)                                      // C1
	nuevoNodo := &NodoMatriz{PosicionX: -1, PosicionY: y, Color: col} // encabezado de las filas
	fila := m.insertarFila(nuevoNodo, m.Raiz)                         // agrega a la fila
	return fila
}

func (m *Matriz) Insertar_Elemento(x int, y int, color string) {
	// agregando posicion en x y y, con color al nuevo nodo
	nuevoNodo := &NodoMatriz{PosicionX: x, PosicionY: y, Color: color}
	nodoColumna := m.buscarColumna(x) // busca la posicion de la columna
	nodoFila := m.buscarFila(y)       //busca la posicion de la fila
	/*
		1. Columna y Fila no Existe
		2. Columna si existe pero Fila no
		3. Fila si existe pero Columna no
		4. Ambos existen
	*/

	if nodoColumna == nil && nodoFila == nil {
		//fmt.Println("Primer Caso")
		nodoColumna = m.nuevaColumna(x)
		nodoFila = m.nuevaFila(y)
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else if nodoColumna != nil && nodoFila == nil {
		//fmt.Println("Segundo Caso")
		nodoFila = m.nuevaFila(y)
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else if nodoColumna == nil && nodoFila != nil {
		//fmt.Println("Tercer Caso")
		nodoColumna = m.nuevaColumna(x)
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else if nodoColumna != nil && nodoFila != nil {
		//fmt.Println("Cuarto Caso")
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else {
		fmt.Println("ERROR!!!!!!")
	}
}

func (m *Matriz) Reporte(nameCapa string) {
	texto := ""
	nombre_archivo := "./" + nameCapa + ".dot"
	nombre_imagen := "./" + nameCapa + ".jpg"
	aux1 := m.Raiz
	aux2 := m.Raiz
	aux3 := m.Raiz
	if aux1 != nil {
		texto = "digraph MatrizCapa{ \n node[shape=box] \n rankdir=UD; \n {rank=min; \n"
		/** Creacion de los nodos actuales */
		for aux1 != nil {
			texto += "nodo" + strconv.Itoa(aux1.PosicionX+1) + strconv.Itoa(aux1.PosicionY+1) + "[label=\"" + aux1.Color + "\" ,rankdir=LR,group=" + strconv.Itoa(aux1.PosicionX+1) + "]; \n"
			aux1 = aux1.Siguiente
		}
		texto += "}"
		for aux2 != nil {
			aux1 = aux2
			texto += "{rank=same; \n"
			for aux1 != nil {
				texto += "nodo" + strconv.Itoa(aux1.PosicionX+1) + strconv.Itoa(aux1.PosicionY+1) + "[label=\"" + aux1.Color + "\" ,group=" + strconv.Itoa(aux1.PosicionX+1) + "]; \n"
				aux1 = aux1.Siguiente
			}
			texto += "}"
			aux2 = aux2.Abajo
		}
		/** Conexiones entre los nodos de la matriz */
		aux2 = aux3
		for aux2 != nil {
			aux1 = aux2
			for aux1.Siguiente != nil {
				texto += "nodo" + strconv.Itoa(aux1.PosicionX+1) + strconv.Itoa(aux1.PosicionY+1) + " -> " + "nodo" + strconv.Itoa(aux1.Siguiente.PosicionX+1) + strconv.Itoa(aux1.Siguiente.PosicionY+1) + " [dir=both];\n"
				aux1 = aux1.Siguiente
			}
			aux2 = aux2.Abajo
		}
		aux2 = aux3
		for aux2 != nil {
			aux1 = aux2
			for aux1.Abajo != nil {
				texto += "nodo" + strconv.Itoa(aux1.PosicionX+1) + strconv.Itoa(aux1.PosicionY+1) + " -> " + "nodo" + strconv.Itoa(aux1.Abajo.PosicionX+1) + strconv.Itoa(aux1.Abajo.PosicionY+1) + " [dir=both];\n"
				aux1 = aux1.Abajo
			}
			aux2 = aux2.Siguiente
		}
		texto += "}"
	} else {
		texto = "No hay elementos en la matriz"
	}
	//fmt.Println(texto)
	crearArchivo(nombre_archivo)
	escribirArchivo(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}

// Funcion para leer la capa
func (m *Matriz) LeerArchivo(ruta string) {

	//listaAux := &ListaCircular{Inicio: nil, Longitud: 0}
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	x := 0
	y := 0
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		for i := 0; i < len(linea); i++ {
			if linea[i] != "x" {
				m.Insertar_Elemento(x, y, linea[i])
			}
			x++
		}
		x = 0
		y++
	}
}

func (m *Matriz) LeerInicial(ruta string, imagen string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		if linea[0] == "0" {
			m.leerConfig("csv/" + imagen + "/" + linea[1]) /*csv/mario/config.csv*/
		} else {
			m.LeerArchivo("csv/" + imagen + "/" + linea[1])
		}
	}
}

func (m *Matriz) LeerInicial1(ruta string, imagen string, listaCapasMatriz *Lista_simpleCapa) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		if linea[0] == "0" {
			m.leerConfig("csv/" + imagen + "/" + linea[1]) /*csv/mario/config.csv*/
		} else {
			//m.LeerArchivo("csv/" + imagen + "/" + linea[1])
			//m.Reporte(linea[1])
			listaCapasMatriz.InsertarCapa(strings.TrimSpace(linea[0]), strings.TrimSpace(linea[1]))
		}
	}
}

func (m *Matriz) LeerInicialYCapaElegida(ruta string, imagen string, capaelegida string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		if linea[0] == "0" {
			m.leerConfig("csv/" + imagen + "/" + linea[1]) /*csv/mario/config.csv*/
		} else {
			if capaelegida == linea[1] {
				m.LeerArchivo("csv/" + imagen + "/" + linea[1])
				m.Reporte(linea[1])
			}
		}
	}
}

func (m *Matriz) leerConfig(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		numero, _ := strconv.Atoi(linea[1])
		if linea[0] == "image_width" {
			m.ImageWidth = numero
		} else if linea[0] == "image_height" {
			m.ImageHeight = numero
		} else if linea[0] == "pixel_width" {
			m.PixelWidth = numero
		} else if linea[0] == "pixel_height" {
			m.PixelHeight = numero
		}
	}
}

func (m *Matriz) GenerarImagen(nombre_imagen string, tipoFiltro string) {
	archivoCSS := "csv/" + nombre_imagen + "/" + nombre_imagen + tipoFiltro + ".css" // csv/mario/mario.css
	contenidoCSS := "body{\n background: #636363; \n background: -webkit-linear-gradient(to right, #636363, #a2ab58);\n background: linear-gradient(to right, #636363, #a2ab58);  \n height: 100vh; \n display: flex; \n justify-content: center; \n align-items: center; \n } \n"
	contenidoCSS += ".canvas{ \n width: " + strconv.Itoa(m.ImageWidth*m.PixelWidth) + "px; \n"
	contenidoCSS += "height: " + strconv.Itoa(m.ImageHeight*m.PixelHeight) + "px; \n }"
	contenidoCSS += ".pixel{ \n width: " + strconv.Itoa(m.PixelWidth) + "px; \n"
	contenidoCSS += "height: " + strconv.Itoa(m.PixelHeight) + "px; \n float: left; \n } \n"
	x_pixel := 0
	x := 1
	auxFila := m.Raiz.Abajo
	auxColumna := auxFila.Siguiente

	for i := 0; i < m.ImageHeight; i++ {
		for j := 0; j < m.ImageWidth; j++ {
			if auxColumna != nil {
				if auxColumna.PosicionX == x_pixel {
					contenidoCSS += ".pixel:nth-child(" + strconv.Itoa(x) + ") { background: rgb(" + strings.ReplaceAll(auxColumna.Color, "-", ",") + "); }\n"
					auxColumna = auxColumna.Siguiente
				}
				x_pixel++
			}
			x++
		}
		x_pixel = 0
		if auxFila.Abajo != nil {
			auxFila = auxFila.Abajo
		}

		if auxFila != nil {
			auxColumna = auxFila.Siguiente
		}
	}

	/*FIN*/
	m.generarHTML(nombre_imagen, tipoFiltro)
	crearArchivo(archivoCSS)
	escribirArchivo(contenidoCSS, archivoCSS)
}

func (m *Matriz) generarHTML(nombre_imagen string, tipoFiltro string) {
	archivoHTML := "csv/" + nombre_imagen + "/" + nombre_imagen + tipoFiltro + ".html"
	contenidoHTML := "<!DOCTYPE html> \n <html> \n <head> \n <link rel=\"stylesheet\"  href=\""
	contenidoHTML += nombre_imagen + tipoFiltro + ".css"
	contenidoHTML += "\" > \n </head> \n <body> \n <div class=\"canvas\"> \n"
	for i := 0; i < m.ImageHeight; i++ {
		for j := 0; j < m.ImageWidth; j++ {
			contenidoHTML += "    <div class=\"pixel\"></div> \n"
		}
	}
	contenidoHTML += "</div> \n </body> \n </html> \n"
	crearArchivo(archivoHTML)
	escribirArchivo(contenidoHTML, archivoHTML)
}

func NewMatriz() *Matriz {
	return &Matriz{nil, 0, 0, 0, 0}
}

// CONFIGURACIONES PARA FILTROS
func (m *Matriz) FiltroNegativo(nombre_imagen string) {
	archivoCSS := "csv/" + nombre_imagen + "/" + nombre_imagen + "Negativo.css" // csv/mario/mario.css
	contenidoCSS := "body{\n background: #636363; \n background: -webkit-linear-gradient(to right, #636363, #a2ab58);\n background: linear-gradient(to right, #636363, #a2ab58);  \n height: 100vh; \n display: flex; \n justify-content: center; \n align-items: center; \n } \n"
	contenidoCSS += ".canvas{ \n width: " + strconv.Itoa(m.ImageWidth*m.PixelWidth) + "px; \n"
	contenidoCSS += "height: " + strconv.Itoa(m.ImageHeight*m.PixelHeight) + "px; \n }"
	contenidoCSS += ".pixel{ \n width: " + strconv.Itoa(m.PixelWidth) + "px; \n"
	contenidoCSS += "height: " + strconv.Itoa(m.PixelHeight) + "px; \n float: left; \n } \n"
	x_pixel := 0
	x := 1
	auxFila := m.Raiz.Abajo
	auxColumna := auxFila.Siguiente

	for i := 0; i < m.ImageHeight; i++ {
		for j := 0; j < m.ImageWidth; j++ {
			if auxColumna != nil {
				if auxColumna.PosicionX == x_pixel {
					//fmt.Println("ANTES DE SUSTITUIR POR LOS NEGATIVOS*************************************************************")
					//fmt.Println(auxColumna.Color)
					//Formato del color 255-51-0     si es 0 se reemplaza por 255 o viceversa o de lo contrario 255-color
					cade := auxColumna.Color
					cad1 := ""
					result := strings.Split(cade, "-")
					for _, elem := range result {
						num := 0
						if elem == "0" {
							num = 255
						} else if elem == "255" {
							num = 0
						} else {
							m, _ := strconv.Atoi(elem)
							num = 255 - m
						}
						cad1 += strconv.Itoa(num) + ","
					}
					nuevoText := strings.TrimSuffix(cad1, string(cad1[len(cad1)-1]))
					auxColumna.Color = strings.ReplaceAll(nuevoText, ",", "-")
					//fmt.Println("DESPUES DE SUSTITUIR POR LOS NEGATIVOS*************************************************************")
					//fmt.Println(auxColumna.Color)
					contenidoCSS += ".pixel:nth-child(" + strconv.Itoa(x) + ") { background: rgb(" + nuevoText + "); }\n"
					auxColumna = auxColumna.Siguiente
				}
				x_pixel++
			}
			x++
		}
		x_pixel = 0
		if auxFila.Abajo != nil {
			auxFila = auxFila.Abajo
		}

		if auxFila != nil {
			auxColumna = auxFila.Siguiente
		}
	}

	/*FIN*/
	m.generarHTMLNegativo(nombre_imagen)
	crearArchivo(archivoCSS)
	escribirArchivo(contenidoCSS, archivoCSS)
}

func (m *Matriz) generarHTMLNegativo(nombre_imagen string) {
	archivoHTML := "csv/" + nombre_imagen + "/" + nombre_imagen + "Negativo.html"
	contenidoHTML := "<!DOCTYPE html> \n <html> \n <head> \n <link rel=\"stylesheet\"  href=\""
	contenidoHTML += nombre_imagen + "Negativo.css"
	contenidoHTML += "\" > \n </head> \n <body> \n <div class=\"canvas\"> \n"
	for i := 0; i < m.ImageHeight; i++ {
		for j := 0; j < m.ImageWidth; j++ {
			contenidoHTML += "    <div class=\"pixel\"></div> \n"
		}
	}
	contenidoHTML += "</div> \n </body> \n </html> \n"
	crearArchivo(archivoHTML)
	escribirArchivo(contenidoHTML, archivoHTML)
}

func (m *Matriz) FiltroEscalaGris(nombre_imagen string) {
	archivoCSS := "csv/" + nombre_imagen + "/" + nombre_imagen + "EscalaGris.css" // csv/mario/mario.css
	contenidoCSS := "body{\n background: #636363; \n background: -webkit-linear-gradient(to right, #636363, #a2ab58);\n background: linear-gradient(to right, #636363, #a2ab58);  \n height: 100vh; \n display: flex; \n justify-content: center; \n align-items: center; \n } \n"
	contenidoCSS += ".canvas{ \n width: " + strconv.Itoa(m.ImageWidth*m.PixelWidth) + "px; \n"
	contenidoCSS += "height: " + strconv.Itoa(m.ImageHeight*m.PixelHeight) + "px; \n }"
	contenidoCSS += ".pixel{ \n width: " + strconv.Itoa(m.PixelWidth) + "px; \n"
	contenidoCSS += "height: " + strconv.Itoa(m.PixelHeight) + "px; \n float: left; \n } \n"
	x_pixel := 0
	x := 1
	auxFila := m.Raiz.Abajo
	auxColumna := auxFila.Siguiente

	for i := 0; i < m.ImageHeight; i++ {
		for j := 0; j < m.ImageWidth; j++ {
			if auxColumna != nil {
				if auxColumna.PosicionX == x_pixel {
					cade := auxColumna.Color
					cad1 := ""
					result := strings.Split(cade, "-")
					num1, _ := strconv.ParseFloat(result[0], 64)
					num2, _ := strconv.ParseFloat(result[1], 64)
					num3, _ := strconv.ParseFloat(result[2], 64)

					escalaGris := 0.299*num1 + 0.587*num2 + 0.114*num3

					cad1 = strconv.Itoa(int(escalaGris)) + ", " + strconv.Itoa(int(escalaGris)) + ", " + strconv.Itoa(int(escalaGris))
					contenidoCSS += ".pixel:nth-child(" + strconv.Itoa(x) + ") { background: rgb(" + cad1 + "); }\n"

					auxColumna.Color = strings.ReplaceAll(cad1, ",", "-")

					auxColumna = auxColumna.Siguiente
				}
				x_pixel++
			}
			x++
		}
		x_pixel = 0
		if auxFila.Abajo != nil {
			auxFila = auxFila.Abajo
		}

		if auxFila != nil {
			auxColumna = auxFila.Siguiente
		}
	}

	/*FIN*/
	m.generarHTMLEscalaGris(nombre_imagen)
	crearArchivo(archivoCSS)
	escribirArchivo(contenidoCSS, archivoCSS)
}

func (m *Matriz) generarHTMLEscalaGris(nombre_imagen string) {
	archivoHTML := "csv/" + nombre_imagen + "/" + nombre_imagen + "EscalaGris.html"
	contenidoHTML := "<!DOCTYPE html> \n <html> \n <head> \n <link rel=\"stylesheet\"  href=\""
	contenidoHTML += nombre_imagen + "EscalaGris.css"
	contenidoHTML += "\" > \n </head> \n <body> \n <div class=\"canvas\"> \n"
	for i := 0; i < m.ImageHeight; i++ {
		for j := 0; j < m.ImageWidth; j++ {
			contenidoHTML += "    <div class=\"pixel\"></div> \n"
		}
	}
	contenidoHTML += "</div> \n </body> \n </html> \n"
	crearArchivo(archivoHTML)
	escribirArchivo(contenidoHTML, archivoHTML)
}

func (m *Matriz) EspejoDoble() {
	var matrizFiltroDoble = &Matriz{Raiz: &NodoMatriz{PosicionX: -1, PosicionY: -1, Color: "RAIZ"}}
	auxFila := m.Raiz.Abajo
	auxColumna := auxFila.Siguiente

	for i := 0; i < m.ImageHeight; i++ {
		for j := 0; j < m.ImageWidth; j++ {
			if auxColumna != nil {
				valorX := (m.ImageWidth - 1) - auxColumna.PosicionX
				valorY := (m.ImageHeight - 1) - auxColumna.PosicionY
				matrizFiltroDoble.Insertar_Elemento(valorX, valorY, auxColumna.Color)
				auxColumna = auxColumna.Siguiente

			}
		}
		if auxFila.Abajo != nil {
			auxFila = auxFila.Abajo
		}

		if auxFila != nil {
			auxColumna = auxFila.Siguiente
		}
	}
	m.Raiz = matrizFiltroDoble.Raiz
}

func (m *Matriz) EspejoX() {
	var matrizFiltroEspX = &Matriz{Raiz: &NodoMatriz{PosicionX: -1, PosicionY: -1, Color: "RAIZ"}}
	auxFila := m.Raiz.Abajo
	auxColumna := auxFila.Siguiente

	for i := 0; i < m.ImageHeight; i++ {
		for j := 0; j < m.ImageWidth; j++ {
			if auxColumna != nil {
				valorX := (m.ImageWidth - 1) - auxColumna.PosicionX
				matrizFiltroEspX.Insertar_Elemento(valorX, auxColumna.PosicionY, auxColumna.Color)
				auxColumna = auxColumna.Siguiente

			}
		}
		if auxFila.Abajo != nil {
			auxFila = auxFila.Abajo
		}

		if auxFila != nil {
			auxColumna = auxFila.Siguiente
		}
	}
	m.Raiz = matrizFiltroEspX.Raiz
}

func (m *Matriz) EspejoY() {
	var matrizFiltroEspY = &Matriz{Raiz: &NodoMatriz{PosicionX: -1, PosicionY: -1, Color: "RAIZ"}}
	auxFila := m.Raiz.Abajo
	auxColumna := auxFila.Siguiente

	for i := 0; i < m.ImageHeight; i++ {
		for j := 0; j < m.ImageWidth; j++ {
			if auxColumna != nil {
				valorY := (m.ImageWidth - 1) - auxColumna.PosicionY
				matrizFiltroEspY.Insertar_Elemento(auxColumna.PosicionX, valorY, auxColumna.Color)
				auxColumna = auxColumna.Siguiente

			}
		}
		if auxFila.Abajo != nil {
			auxFila = auxFila.Abajo
		}

		if auxFila != nil {
			auxColumna = auxFila.Siguiente
		}
	}
	m.Raiz = matrizFiltroEspY.Raiz
}
