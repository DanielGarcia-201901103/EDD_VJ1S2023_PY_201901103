package estructura

import (
	"fmt"
)

type Pedido struct {
	idCliente  string
	idEmpleado string
	nameImagen string
	idImagen   string
}

type NodoPila struct {
	data      *Pedido
	siguiente *NodoPila
}

type Pila struct {
	Primero  *NodoPila
	Longitud int
}

func (pila *Pila) Push(idCliente string, idEmpleado string, nameImagen string, idImagen string) {
	nuevoPedido := &Pedido{idCliente: idCliente, idEmpleado: idEmpleado, nameImagen: nameImagen, idImagen: idImagen}
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
		text += "|(ID Cliente: " + aux.data.idCliente + "\\" + "n Imagen:" + aux.data.nameImagen +
			")"
		aux = aux.siguiente
	}
	text += "\"]; \n}"
	crearArchivo(nombreArchivo)
	escribirArchivo(text, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}
