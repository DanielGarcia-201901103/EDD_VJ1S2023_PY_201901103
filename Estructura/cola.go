package estructura

import "fmt"

type ClienteCola struct {
	id   string
	name string
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
	nuevoCliente := &ClienteCola{id: id, name: name}
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
		cola.Longitud++
	}
}
