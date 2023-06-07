package estructura

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
		lista.Longitud++
	} else {
		aux := lista.Inicio
		for aux.siguiente != nil { //esto simula un while
			aux = aux.siguiente
		}
		aux.siguiente = &Nodo_circular{data: cliente, siguiente: lista.Inicio}
		lista.Longitud++
	}
}
