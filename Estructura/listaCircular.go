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
