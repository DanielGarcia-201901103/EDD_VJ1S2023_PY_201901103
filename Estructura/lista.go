package estructura

import "fmt"

type Lista struct {
	Inicio   *Nodo
	Longitud int
}

func (lista *Lista) estaVacia() bool {
	if lista.Longitud == 0 {
		return true
	}
	return false
}

func (lista *Lista) Insertar(numero int) {
	if lista.estaVacia() {
		lista.Inicio = &Nodo{valor: numero, siguiente: nil}
		lista.Longitud++
	} else {
		aux := lista.Inicio
		for aux.siguiente != nil { //esto simula un while
			aux = aux.siguiente
		}
		aux.siguiente = &Nodo{valor: numero, siguiente: nil}
		lista.Longitud++
	}
}

func (lista *Lista) Mostrar() {
	aux := lista.Inicio

	for aux != nil {
		fmt.Println(aux.valor)
		aux = aux.siguiente
	}
}
