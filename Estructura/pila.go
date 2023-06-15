package estructura

import (
	"fmt"
	"strings"
)

type Pedido struct {
	idCliente  string
	idEmpleado string
	nameImagen string
}

type NodoPila struct {
	data      *Pedido
	siguiente *NodoPila
}

type Pila struct {
	Primero  *NodoPila
	Longitud int
}

func (pila *Pila) Push(idCliente string, idEmpleado string, nameImagen string) {
	nuevoPedido := &Pedido{idCliente: idCliente, idEmpleado: idEmpleado, nameImagen: nameImagen}
	if pila.Longitud == 0 {
		nuevoN := &NodoPila{data: nuevoPedido, siguiente: nil}
		pila.Primero = nuevoN
		pila.Longitud++
	} else {
		nuevoN := &NodoPila{data: nuevoPedido, siguiente: pila.Primero}
		pila.Primero = nuevoN
		pila.Longitud++
	}
}

func (pila *Pila) Pop() {
	if pila.Longitud == 0 {
		fmt.Println("No hay elementos en la pila")
	} else {
		pila.Primero = pila.Primero.siguiente
		pila.Longitud--
	}
}

func (pila *Pila) ReportePila() {
	nombreArchivo := "./pila.dot"
	nombreImagen := "./pila.jpg"
	text := "digraph pila{\n"
	text += "rankdir = LR; \n"
	text += "node[shape = record]; \n"
	aux := pila.Primero
	text += "nodo0 [label=\""
	//text += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < pila.Longitud; i++ {
		text += "|(ID Cliente: " + aux.data.idCliente + "\\" + "n Imagen: " + aux.data.nameImagen +
			")"
		aux = aux.siguiente
	}
	text += "\"]; \n}"
	crearArchivo(nombreArchivo)
	escribirArchivo(text, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}

func (pila *Pila) ReporteJson() {
	nombreArchivo := "./pedidos.json"
	text := "{\n"
	text += "\t\"pedidos\":["
	aux := pila.Primero
	for i := 0; i < pila.Longitud; i++ {
		text += "\n\t\t{\n"
		text += "\t\t\t\"id_cliente\": " + aux.data.idCliente + ",\n"
		text += "\t\t\t\"imagen\": \"" + aux.data.nameImagen + "\"\n"
		text += "\t\t},"
		aux = aux.siguiente
	}
	nuevoText := strings.TrimSuffix(text, string(text[len(text)-1]))
	nuevoText += "\n\t]\n}"

	crearArchivo(nombreArchivo)
	escribirArchivo(nuevoText, nombreArchivo)
}

func NewPila() *Pila {
	return &Pila{nil, 0}
}
