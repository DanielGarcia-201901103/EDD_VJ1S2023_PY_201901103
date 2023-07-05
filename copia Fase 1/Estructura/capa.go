package estructura

import (
	"fmt"
	"strconv"
)

type Capa struct {
	layer_capa string
	file_capa  string
}

type NodoCapa struct {
	data      *Capa
	siguiente *NodoCapa
}

type Lista_simpleCapa struct {
	Inicio   *NodoCapa
	Longitud int
}

func (lista *Lista_simpleCapa) estaVacia() bool {
	if lista.Longitud == 0 {
		return true
	}
	return false
}

// Inserta al final
func (lista *Lista_simpleCapa) InsertarCapa(layer string, file string) {
	capas := &Capa{layer_capa: layer, file_capa: file}
	if lista.estaVacia() {
		lista.Inicio = &NodoCapa{data: capas, siguiente: nil}
		lista.Longitud++
	} else {
		aux := lista.Inicio
		for aux.siguiente != nil { //esto simula un while
			aux = aux.siguiente
		}
		aux.siguiente = &NodoCapa{data: capas, siguiente: nil}
		lista.Longitud++
	}
}

func (lista *Lista_simpleCapa) ListarDatosCapa() {
	aux := lista.Inicio
	var contador int = 1
	for aux != nil {
		fmt.Println(strconv.Itoa(contador)+". ", aux.data.file_capa, "capa: ", aux.data.layer_capa)
		aux = aux.siguiente
		contador++
	}
}

func (lista *Lista_simpleCapa) BuscarCapa(opcion string) string {
	aux := lista.Inicio
	var contador int = 1
	for aux != nil {
		if opcion == strconv.Itoa(contador) {
			return aux.data.file_capa
		}
		//fmt.Println(strconv.Itoa(contador)+". ", aux.data.name_Imagen)
		aux = aux.siguiente
		contador++
	}
	return "No se encuentra la opción elegida"
}

func NewListaSimpleCapa() *Lista_simpleCapa {
	return &Lista_simpleCapa{nil, 0}
}

/*
Agregar nodo a la lista, al final o al inicio
buscar nodo de la lista
actualizar nodo de la lista, primero busca y luego actualiza
borrar nodo de la lista,  primero busca y luego elimina*/
