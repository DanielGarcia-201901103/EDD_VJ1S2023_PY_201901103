package estructura

import (
	"fmt"
	"strconv"
)

type Cliente struct {
	id_cliente   string
	name_cliente string
}

type Nodo_circular struct {
	data      *Cliente
	siguiente *Nodo_circular
}

type Lista_CircularS struct {
	Inicio   *Nodo_circular
	Longitud int
}

func (lista *Lista_CircularS) estaVacia() bool {
	if lista.Longitud == 0 {
		return true
	}
	return false
}

// Inserta al final
func (lista *Lista_CircularS) Insertar(id_cliente string, name_cliente string) {
	cliente := &Cliente{id_cliente: id_cliente, name_cliente: name_cliente}
	if lista.estaVacia() {
		lista.Inicio = &Nodo_circular{data: cliente, siguiente: nil}
		lista.Inicio.siguiente = lista.Inicio
		lista.Longitud++
	} else {
		if lista.Longitud == 1 {
			lista.Inicio.siguiente = &Nodo_circular{data: cliente, siguiente: lista.Inicio}
			lista.Longitud++
		} else {
			aux := lista.Inicio
			for iterar := 0; iterar < lista.Longitud-1; iterar++ {
				aux = aux.siguiente
			}
			aux.siguiente = &Nodo_circular{data: cliente, siguiente: lista.Inicio}
			lista.Longitud++
		}
	}
}

func (lista *Lista_CircularS) Mostrar() {
	aux := lista.Inicio
	for iterar := 0; iterar < lista.Longitud; iterar++ {
		fmt.Println(aux.data.id_cliente, " ", aux.data.name_cliente)
		aux = aux.siguiente
	}
}

func (lista *Lista_CircularS) ReporteCircular() {
	nombreArchivo := "./listaCircularSimple.dot"
	nombreImagen := "./listadoCircularSimple.jpg"
	text := "digraph lista{\n"
	text += "rankdir = LR; \n"
	text += "node[shape = record]; \n"
	//text += "nodonull1[label=\"null\"];\n"
	//text += "nodonull2[label=\"null\"];\n"
	aux := lista.Inicio
	contador := 0
	//text += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < lista.Longitud; i++ {
		text += "nodo" + strconv.Itoa(i) + "[label =\" ID:" + aux.data.id_cliente + "\\" + "n Nombre: " + aux.data.name_cliente + "\"]; \n"
		aux = aux.siguiente
	}

	for i := 0; i < lista.Longitud-1; i++ {
		c := i + 1
		text += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		//text += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	text += "nodo" + strconv.Itoa(contador) + "->nodo0;\n"
	text += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(text, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}

func NewListaCircular() *Lista_CircularS {
	return &Lista_CircularS{nil, 0}
}
