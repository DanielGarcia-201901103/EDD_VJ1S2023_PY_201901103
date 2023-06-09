package estructura

type Pedido struct {
	idCliente  string
	idEmpleado string
	nameImagen string
	idImagen   string
}

type NodoPila struct {
	data      *Pedido
	siguiente *NodoPila
}

type Pila struct {
	Primero  *NodoPila
	Longitud int
}

func (pila *Pila) Push(idCliente string, idEmpleado string, nameImagen string, idImagen string) {
	nuevoPedido := &Pedido{idCliente: idCliente, idEmpleado: idEmpleado, nameImagen: nameImagen, idImagen: idImagen}
	if pila.Longitud == 0 {
		nuevoN := &NodoPila{data: nuevoPedido, siguiente: nil}
		pila.Primero = nuevoN
		pila.Longitud++
	}
}
