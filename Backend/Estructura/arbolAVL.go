package estructura

import (
	"math"
	"strconv"
)

type NodoAVL struct {
	Izquierdo      *NodoAVL
	Derecho        *NodoAVL
	Data           int //valor
	Altura         int
	EquilibrioFact int
}

type ArbolAVL struct {
	Raiz *NodoAVL
}

func (arbolAVL *ArbolAVL) altura(raiz *NodoAVL) int {
	if raiz == nil {
		return 0
	}
	return raiz.Altura
}

func (arbolAVL *ArbolAVL) equilibrio(raiz *NodoAVL) int {
	if raiz == nil {
		return 0
	}
	return (arbolAVL.altura(raiz.Derecho) - arbolAVL.altura(raiz.Izquierdo))
}

func (arbolAVL *ArbolAVL) InsertarElemento(data int) {
	newNode := &NodoAVL{Data: data}
	arbolAVL.Raiz = arbolAVL.insertarNodo(arbolAVL.Raiz, newNode)
}

func (arbolAVL *ArbolAVL) insertarNodo(raiz *NodoAVL, newNodo *NodoAVL) *NodoAVL {
	if raiz == nil {
		raiz = newNodo
	} else {
		if raiz.Data > newNodo.Data {
			raiz.Izquierdo = arbolAVL.insertarNodo(raiz.Izquierdo, newNodo)
		} else {
			raiz.Derecho = arbolAVL.insertarNodo(raiz.Derecho, newNodo)
		}
	}
	numMaximo := math.Max(float64(arbolAVL.altura(raiz.Izquierdo)), float64(arbolAVL.altura(raiz.Derecho)))
	raiz.Altura = 1 + int(numMaximo)
	balanceando := arbolAVL.equilibrio(raiz)
	raiz.EquilibrioFact = balanceando
	if balanceando > 1 && newNodo.Data > raiz.Derecho.Data {
		return arbolAVL.rotIzquierda(raiz)
	}
	if balanceando < -1 && newNodo.Data < raiz.Izquierdo.Data {
		return arbolAVL.rotDerecha(raiz)
	}
	if balanceando > 1 && newNodo.Data < raiz.Derecho.Data {
		raiz.Derecho = arbolAVL.rotDerecha(raiz.Derecho)
		return arbolAVL.rotIzquierda(raiz)
	}
	if balanceando < -1 && newNodo.Data > raiz.Izquierdo.Data {
		raiz.Izquierdo = arbolAVL.rotIzquierda(raiz.Izquierdo)
		return arbolAVL.rotDerecha(raiz)
	}
	return raiz
}

func (arbolAVL *ArbolAVL) rotIzquierda(raiz *NodoAVL) *NodoAVL {
	raizDerecha := raiz.Derecho
	childIzquierdo := raizDerecha.Izquierdo
	raizDerecha.Izquierdo = raiz
	raiz.Derecho = childIzquierdo
	numMax := math.Max(float64(arbolAVL.altura(raiz.Izquierdo)), float64(arbolAVL.altura(raiz.Derecho)))
	raiz.Altura = 1 + int(numMax)
	numMax = math.Max(float64(arbolAVL.altura(raizDerecha.Izquierdo)), float64(arbolAVL.altura(raizDerecha.Derecho)))
	raizDerecha.Altura = 1 + int(numMax)
	raiz.EquilibrioFact = arbolAVL.equilibrio(raiz)
	raizDerecha.EquilibrioFact = arbolAVL.equilibrio(raizDerecha)
	return raizDerecha
}

func (arbolAVL *ArbolAVL) rotDerecha(raiz *NodoAVL) *NodoAVL {
	raizIzquierda := raiz.Izquierdo
	childDerecha := raizIzquierda.Derecho
	raizIzquierda.Derecho = raiz
	raiz.Izquierdo = childDerecha
	numMax := math.Max(float64(arbolAVL.altura(raiz.Izquierdo)), float64(arbolAVL.altura(raiz.Derecho)))
	raiz.Altura = 1 + int(numMax)
	numMax = math.Max(float64(arbolAVL.altura(raizIzquierda.Izquierdo)), float64(arbolAVL.altura(raizIzquierda.Derecho)))
	raizIzquierda.Altura = 1 + int(numMax)
	raiz.EquilibrioFact = arbolAVL.equilibrio(raiz)
	raizIzquierda.EquilibrioFact = arbolAVL.equilibrio(raizIzquierda)
	return raizIzquierda
}

func (arbolAVL *ArbolAVL) Graficar() {
	text := ""
	nameArchivo := "./arbolAVL.dot"
	nameImagen := "arbolAVL.jpg"
	if arbolAVL.Raiz != nil {
		text += "digraph arbolAVL{"
		text += arbolAVL.valArbol(arbolAVL.Raiz, 0)
		text += "}"
	}
	crearArchivo(nameArchivo)
	escribirArchivo(text, nameArchivo)
	ejecutar(nameImagen, nameArchivo)
}

func (a *ArbolAVL) valArbol(raiz *NodoAVL, indice int) string {
	text := ""
	indice1 := indice + 1
	if raiz != nil {
		text += "\""
		text += strconv.Itoa(raiz.Data)
		text += "\" ;"
		if raiz.Izquierdo != nil && raiz.Derecho != nil {
			text += " x" + strconv.Itoa(indice1) + " [label=\"\",width=.1,style=invis];"
			text += "\""
			text += strconv.Itoa(raiz.Data)
			text += "\" -> "
			text += a.valArbol(raiz.Izquierdo, indice1)
			text += "\""
			text += strconv.Itoa(raiz.Data)
			text += "\" -> "
			text += a.valArbol(raiz.Derecho, indice1)
			text += "{rank=same" + "\"" + strconv.Itoa(raiz.Izquierdo.Data) + "\"" + " -> " + "\"" + strconv.Itoa(raiz.Derecho.Data) + "\"" + " [style=invis]}; "
		} else if raiz.Izquierdo != nil && raiz.Derecho == nil {
			text += " x" + strconv.Itoa(indice1) + " [label=\"\",width=.1,style=invis];"
			text += "\""
			text += strconv.Itoa(raiz.Data)
			text += "\" -> "
			text += a.valArbol(raiz.Izquierdo, indice1)
			text += "\""
			text += strconv.Itoa(raiz.Data)
			text += "\" -> "
			text += "x" + strconv.Itoa(indice1) + "[style=invis]"
			text += "{rank=same" + "\"" + strconv.Itoa(raiz.Izquierdo.Data) + "\"" + " -> " + "x" + strconv.Itoa(indice1) + " [style=invis]}; "
		} else if raiz.Izquierdo == nil && raiz.Derecho != nil {
			text += " x" + strconv.Itoa(indice1) + " [label=\"\",width=.1,style=invis];"
			text += "\""
			text += strconv.Itoa(raiz.Data)
			text += "\" -> "
			text += "x" + strconv.Itoa(indice1) + "[style=invis]"
			text += "; \""
			text += strconv.Itoa(raiz.Data)
			text += "\" -> "
			text += a.valArbol(raiz.Derecho, indice1)
			text += "{rank=same" + " x" + strconv.Itoa(indice1) + " -> \"" + strconv.Itoa(raiz.Derecho.Data) + "\"" + " [style=invis]}; "
		}
	}
	return text
}
