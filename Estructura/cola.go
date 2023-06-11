package estructura

import (
	"fmt"
	"strconv"
)

type ClienteCola struct {
	id_cola   string
	name_cola string
}

type NodoCola struct {
	data      *ClienteCola
	siguiente *NodoCola
}

type Cola struct {
	Primero  *NodoCola
	Longitud int
}

func (cola *Cola) Encolar(id string, name string) {
	nuevoCliente := &ClienteCola{id_cola: id, name_cola: name}
	if cola.Longitud == 0 {
		nuevoN := &NodoCola{nuevoCliente, nil}
		cola.Primero = nuevoN
		cola.Longitud++
	} else {
		nuevoN := &NodoCola{nuevoCliente, nil}
		aux := cola.Primero
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = nuevoN
		cola.Longitud++
	}
}

func (cola *Cola) Descolar() {
	if cola.Longitud == 0 {
		fmt.Println("Cola vacía")
	} else {
		cola.Primero = cola.Primero.siguiente
		cola.Longitud--
	}
}

func (cola *Cola) ReporteCola() {
	nombreArchivo := "./cola.dot"
	nombreImagen := "./cola.jpg"
	text := "digraph cola{\n"
	text += "rankdir = LR; \n"
	text += "node[shape = record]; \n"
	text += "nodonull1[label=\"null\"];\n"
	//text += "nodonull2[label=\"null\"];\n"
	aux := cola.Primero
	contador := 0
	//text += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < cola.Longitud; i++ {
		text += "nodo" + strconv.Itoa(i) + "[label=\"{ID: " + aux.data.id_cola + "\\" + "n Nombre: " + aux.data.name_cola + "|}\"];\n"
		aux = aux.siguiente
	}

	for i := 0; i < cola.Longitud-1; i++ {
		c := i + 1
		text += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		//text += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	text += "nodo" + strconv.Itoa(contador) + "->nodonull1;\n"
	text += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(text, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}

func NewCola() *Cola {
	return &Cola{nil, 0}
}
