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
