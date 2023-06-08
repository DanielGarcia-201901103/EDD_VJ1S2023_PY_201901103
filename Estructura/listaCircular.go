package estructura

import "fmt"

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

func NewListaCircular() *Lista_CircularS {
	return &Lista_CircularS{nil, 0}
}
