package estructura

import (
	"fmt"
	"strconv"
)

type Imagen struct {
	name_Imagen    string
	cantidad_Capas string
}

type Nodo_Doble struct {
	data      *Imagen
	siguiente *Nodo_Doble
	anterior  *Nodo_Doble
}

type Lista_doble struct {
	Inicio   *Nodo_Doble
	Final    *Nodo_Doble
	Longitud int
}

func (lista *Lista_doble) estaVacia() bool {
	if lista.Longitud == 0 {
		return true
	}
	return false
}

// Inserta al final
func (lista *Lista_doble) Insertar(name_Imagen string, cantidad_Capas string) {
	imagen_c := &Imagen{name_Imagen: name_Imagen, cantidad_Capas: cantidad_Capas}

	if lista.estaVacia() {
		lista.Inicio = &Nodo_Doble{data: imagen_c, siguiente: nil, anterior: nil}
		lista.Final = &Nodo_Doble{data: imagen_c, siguiente: nil, anterior: nil}
		lista.Longitud++
	} else {
		aux := lista.Inicio
		for aux.siguiente != nil { //esto simula un while
			aux = aux.siguiente
		}
		aux.siguiente = &Nodo_Doble{data: imagen_c, siguiente: nil, anterior: aux}
		//aux.siguiente.anterior = aux  esto va si en anterior va nil
		lista.Longitud++
	}
}

func (lista *Lista_doble) MostrarAscendente() {
	aux := lista.Inicio
	for aux != nil {
		fmt.Print(aux.data.name_Imagen)
		fmt.Println(" --> ", aux.data.cantidad_Capas)
		aux = aux.siguiente
	}
}

func (lista *Lista_doble) MostrarDescendente() {
	aux := lista.Inicio
	for aux.siguiente != nil {
		//fmt.Print(aux.data.name_Imagen)
		//fmt.Println(" --> ", aux.data.cantidad_Capas)
		aux = aux.siguiente
	}
	for aux != nil {
		fmt.Print(aux.data.name_Imagen)
		fmt.Println(" --> ", aux.data.cantidad_Capas)
		aux = aux.anterior
	}
}

func (lista *Lista_doble) ListarDatos() {
	aux := lista.Inicio
	var contador int = 1
	for aux != nil {
		fmt.Println(strconv.Itoa(contador)+". ", aux.data.name_Imagen)
		aux = aux.siguiente
		contador++
	}
}

func (lista *Lista_doble) BuscarImagen(opcion string) string {
	aux := lista.Inicio
	var contador int = 1
	for aux != nil {
		if opcion == strconv.Itoa(contador) {
			return aux.data.name_Imagen
		}
		//fmt.Println(strconv.Itoa(contador)+". ", aux.data.name_Imagen)
		aux = aux.siguiente
		contador++
	}
	return "No se encuentra la opción elegida"
}

func (lista *Lista_doble) ReporteDoble() {
	nombreArchivo := "./listadoble.dot"
	nombreImagen := "./listadoble.jpg"
	text := "digraph listaDoble{\n"
	text += "rankdir = LR; \n"
	text += "node[shape = record]; \n"
	text += "nodonull1[label=\"null\"];\n"
	text += "nodonull2[label=\"null\"];\n"
	aux := lista.Inicio
	contador := 0
	text += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < lista.Longitud; i++ {
		text += "nodo" + strconv.Itoa(i) + "[label =\" " + aux.data.name_Imagen + "\"]; \n"
		aux = aux.siguiente
	}
	for i := 0; i < lista.Longitud-1; i++ {
		c := i + 1
		text += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		text += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	text += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	text += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(text, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}

func NewListaDoble() *Lista_doble {
	return &Lista_doble{nil, nil, 0}
}
