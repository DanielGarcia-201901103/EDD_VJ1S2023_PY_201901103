package estructura

import (
	"fmt"
	"strconv"
)

type Empleado struct {
	id     string
	name   string
	cargo  string
	passwd string
}

type Nodo struct {
	data      *Empleado
	siguiente *Nodo
}

type Lista_simple struct {
	Inicio   *Nodo
	Longitud int
}

func (lista *Lista_simple) estaVacia() bool {
	if lista.Longitud == 0 {
		return true
	}
	return false
}

// Inserta al final
func (lista *Lista_simple) Insertar(id string, name string, cargo string, passwd string) {
	empleado := &Empleado{id: id, name: name, cargo: cargo, passwd: passwd}
	if lista.estaVacia() {
		lista.Inicio = &Nodo{data: empleado, siguiente: nil}
		lista.Longitud++
	} else {
		aux := lista.Inicio
		for aux.siguiente != nil { //esto simula un while
			aux = aux.siguiente
		}
		aux.siguiente = &Nodo{data: empleado, siguiente: nil}
		lista.Longitud++
	}
}
func (lista *Lista_simple) Validar(usuario string, password string) bool {
	aux := lista.Inicio
	for aux != nil {
		if usuario == aux.data.id && password == aux.data.passwd {
			return true
		}
		aux = aux.siguiente
	}
	return false
}

func (lista *Lista_simple) Mostrar() {
	aux := lista.Inicio

	for aux != nil {
		fmt.Println(aux.data.id, " ", aux.data.name, " ", aux.data.cargo, " ", aux.data.passwd)
		aux = aux.siguiente
	}
}

func (lista *Lista_simple) ReporteSimple() {
	nombreArchivo := "./listasimple.dot"
	nombreImagen := "./listadosimple.jpg"
	text := "digraph listaSimple{\n"
	text += "rankdir = LR; \n"
	text += "node[shape = record]; \n"
	//text += "nodonull1[label=\"null\"];\n"
	//text += "nodonull2[label=\"null\"];\n"
	aux := lista.Inicio
	//contador := 0
	//text += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < lista.Longitud; i++ {
		text += "nodo" + strconv.Itoa(i) + "[label =\" ID:" + aux.data.id + "\\" + "n Nombre: " + aux.data.name + "\"]; \n"
		aux = aux.siguiente
	}

	for i := 0; i < lista.Longitud-1; i++ {
		c := i + 1
		text += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		//text += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
	}
	//text += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	text += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(text, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}

func NewListaSimple() *Lista_simple {
	return &Lista_simple{nil, 0}
}

/*
Agregar nodo a la lista, al final o al inicio
buscar nodo de la lista
actualizar nodo de la lista, primero busca y luego actualiza
borrar nodo de la lista,  primero busca y luego elimina*/
