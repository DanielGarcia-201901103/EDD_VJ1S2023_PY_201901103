package estructura

import "fmt"

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

func (lista *Lista_simple) Mostrar() {
	aux := lista.Inicio

	for aux != nil {
		fmt.Println(aux.data.id, " ", aux.data.name, " ", aux.data.cargo, " ", aux.data.passwd)
		aux = aux.siguiente
	}
}
