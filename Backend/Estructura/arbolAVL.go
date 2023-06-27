package estructura

import (
	"math"
	"strconv"
)

type PedidosAVL struct {
	Id_Cliente     int
	imagen_Cliente string
}

type NodoAVL struct {
	Izquierdo      *NodoAVL
	Derecho        *NodoAVL
	Data           *PedidosAVL
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

func (arbolAVL *ArbolAVL) InsertarElemento(id int, imagen string) {
	newPedido := &PedidosAVL{Id_Cliente: id, imagen_Cliente: imagen}
	newNode := &NodoAVL{Data: newPedido}
	arbolAVL.Raiz = arbolAVL.insertarNodo(arbolAVL.Raiz, newNode)
}

func (arbolAVL *ArbolAVL) insertarNodo(raiz *NodoAVL, newNode *NodoAVL) *NodoAVL {
	if raiz == nil {
		raiz = newNode
	} else {
		if raiz.Data.Id_Cliente > newNode.Data.Id_Cliente {
			raiz.Izquierdo = arbolAVL.insertarNodo(raiz.Izquierdo, newNode)
		} else {
			raiz.Derecho = arbolAVL.insertarNodo(raiz.Derecho, newNode)
		}
	}
	numMaximo := math.Max(float64(arbolAVL.altura(raiz.Izquierdo)), float64(arbolAVL.altura(raiz.Derecho)))
	raiz.Altura = 1 + int(numMaximo)
	balanceando := arbolAVL.equilibrio(raiz)
	raiz.EquilibrioFact = balanceando
	if balanceando > 1 && newNode.Data.Id_Cliente > raiz.Derecho.Data.Id_Cliente {
		return arbolAVL.rotIzquierda(raiz)
	}
	if balanceando < -1 && newNode.Data.Id_Cliente < raiz.Izquierdo.Data.Id_Cliente {
		return arbolAVL.rotDerecha(raiz)
	}
	if balanceando > 1 && newNode.Data.Id_Cliente < raiz.Derecho.Data.Id_Cliente {
		raiz.Derecho = arbolAVL.rotDerecha(raiz.Derecho)
		return arbolAVL.rotIzquierda(raiz)
	}
	if balanceando < -1 && newNode.Data.Id_Cliente > raiz.Izquierdo.Data.Id_Cliente {
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
	numero := indice + 1
	if raiz != nil {
		text += "\""
		text += strconv.Itoa(raiz.Data.Id_Cliente) + " - " + raiz.Data.imagen_Cliente
		text += "\" ;"
		if raiz.Izquierdo != nil && raiz.Derecho != nil {
			text += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			text += "\""
			text += strconv.Itoa(raiz.Data.Id_Cliente) + " - " + raiz.Data.imagen_Cliente
			text += "\" -> "
			text += a.valArbol(raiz.Izquierdo, numero)
			text += "\""
			text += strconv.Itoa(raiz.Data.Id_Cliente) + " - " + raiz.Data.imagen_Cliente
			text += "\" -> "
			text += a.valArbol(raiz.Derecho, numero)
			text += "{rank=same" + "\"" + strconv.Itoa(raiz.Izquierdo.Data.Id_Cliente) + " - " + raiz.Izquierdo.Data.imagen_Cliente + "\"" + " -> " + "\"" + strconv.Itoa(raiz.Derecho.Data.Id_Cliente) + " - " + raiz.Derecho.Data.imagen_Cliente + "\"" + " [style=invis]}; "
		} else if raiz.Izquierdo != nil && raiz.Derecho == nil {
			text += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			text += "\""
			text += strconv.Itoa(raiz.Data.Id_Cliente) + " - " + raiz.Data.imagen_Cliente
			text += "\" -> "
			text += a.valArbol(raiz.Izquierdo, numero)
			text += "\""
			text += strconv.Itoa(raiz.Data.Id_Cliente) + " - " + raiz.Data.imagen_Cliente
			text += "\" -> "
			text += "x" + strconv.Itoa(numero) + "[style=invis]"
			text += "{rank=same" + "\"" + strconv.Itoa(raiz.Izquierdo.Data.Id_Cliente) + " - " + raiz.Izquierdo.Data.imagen_Cliente + "\"" + " -> " + "x" + strconv.Itoa(numero) + " [style=invis]}; "
		} else if raiz.Izquierdo == nil && raiz.Derecho != nil {
			text += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			text += "\""
			text += strconv.Itoa(raiz.Data.Id_Cliente) + " - " + raiz.Data.imagen_Cliente
			text += "\" -> "
			text += "x" + strconv.Itoa(numero) + "[style=invis]"
			text += "; \""
			text += strconv.Itoa(raiz.Data.Id_Cliente) + " - " + raiz.Data.imagen_Cliente
			text += "\" -> "
			text += a.valArbol(raiz.Derecho, numero)
			text += "{rank=same" + " x" + strconv.Itoa(numero) + " -> \"" + strconv.Itoa(raiz.Derecho.Data.Id_Cliente) + " - " + raiz.Derecho.Data.imagen_Cliente + "\"" + " [style=invis]}; "
		}
	}
	return text
}

func (a *ArbolAVL) InOrder(clientesCola *Cola) {
	a.inOrder1(a.Raiz, clientesCola)
}

func (a *ArbolAVL) inOrder1(tmp *NodoAVL, clientesCola *Cola) {
	if tmp != nil {
		a.inOrder1(tmp.Izquierdo, clientesCola)
		//fmt.Println(tmp.Data.Id_Cliente)
		clientesCola.Encolar(strconv.Itoa(tmp.Data.Id_Cliente), tmp.Data.imagen_Cliente)
		a.inOrder1(tmp.Derecho, clientesCola)
	}
}
