package estructura

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

type NodoBlock struct {
	Bloque    map[string]string
	Siguiente *NodoBlock
	Anterior  *NodoBlock
}

type NodoBlockPet struct {
	Timestamp string
	Biller    string
	Customer  string
	Payment   string
}

type RespBlock struct {
	Id      string
	Factura string
}

type BlockChain struct {
	Inicio          *NodoBlock
	Bloques_Creados int
}

func (b *BlockChain) InsertarBloque(fecha string, biller string, customer string, payment string) {
	cadenaFuncion := strconv.Itoa(b.Bloques_Creados) + fecha + biller + customer + payment
	hash := SHA256(cadenaFuncion)
	if b.Bloques_Creados == 0 {
		datosBloque := map[string]string{
			"index":        strconv.Itoa(b.Bloques_Creados),
			"timestamp":    fecha,
			"biller":       biller,
			"customer":     customer,
			"payment":      payment,
			"previoushash": "0000",
			"hash":         hash,
		}
		nuevoBloque := &NodoBlock{Bloque: datosBloque}
		b.Inicio = nuevoBloque
	} else {
		aux := b.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		datosBloque := map[string]string{
			"index":        strconv.Itoa(b.Bloques_Creados),
			"timestamp":    fecha,
			"biller":       biller,
			"customer":     customer,
			"payment":      payment,
			"previoushash": aux.Bloque["hash"],
			"hash":         hash,
		}
		nuevoBloque := &NodoBlock{Bloque: datosBloque, Anterior: aux}
		aux.Siguiente = nuevoBloque
	}
	b.Bloques_Creados++
}

func SHA256(cadena string) string {
	hexaString := ""
	h := sha256.New()
	h.Write([]byte(cadena))
	hash := h.Sum(nil)
	hexaString = hex.EncodeToString(hash)
	return hexaString
}

func (b *BlockChain) ArregloFacturas() []RespBlock {
	aux := b.Inicio
	var finalArreglo []RespBlock
	for aux != nil {
		finalArreglo = append(finalArreglo, RespBlock{Id: aux.Bloque["customer"], Factura: aux.Bloque["hash"]})
		aux = aux.Siguiente
	}
	fmt.Println(finalArreglo)
	return finalArreglo
}

func (b *BlockChain) InsertTabla(tabla *TablaHash, idEmpleado string) {
	aux := b.Inicio
	for aux != nil {
		if aux.Bloque["biller"] == idEmpleado {
			tabla.Insertar(aux.Bloque["customer"], aux.Bloque["hash"])
		}
		aux = aux.Siguiente
	}
}

func (b *BlockChain) ReporteBloque() {
	cadena := "digraph Bloque{ \n node [margin=0 fontcolor=black fontsize=25 shape=rectangle color=bisque3 style=filled margin = 0.3];\n"
	nombre_archivo := "./bloquePagos.dot"
	nombre_imagen := "./bloquePagos.jpg"
	aux := b.Inicio
	i := 0
	longitud := 0
	for aux != nil {
		//"TimeStamp: 01-06-2023-::16:05:42 \n Biller: 2566 \nCustomer: 9536\nPreviousHash: 0000"
		cadena += "nodo" + strconv.Itoa(i) + "[label=\"TimeStamp: " + aux.Bloque["timestamp"] + "\\" + "n Biller: " + aux.Bloque["biller"] + "\\" + "n Customer: " + aux.Bloque["customer"] + "\\" + "n PreviousHash: " + aux.Bloque["previoushash"] + "\"]; \n"
		i++
		longitud++
		aux = aux.Siguiente
	}
	i = 0
	for i := 0; i < longitud-1; i++ {
		c := i + 1
		cadena += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
	}
	cadena += "\n}"
	crearArchivo(nombre_archivo)
	escribirArchivo(cadena, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}
