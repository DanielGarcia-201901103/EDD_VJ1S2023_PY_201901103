# EDD_VJ1S2023_PY_201901103

## **MANUAL USUARIO**

EDD Creative es un proyecto que permite la interacción con el usuario a traves consola, contando con el menú principal, en el cual el administrador es el que carga los datos de los empleados, imagenes, y clientes, por medio de archivos csv.

> #### Login
>
> Cuando se inicia el programa solamente podrá acceder al menú administrador con el siguiente usuario: ADMIN_201901103 y con la contraseña: Admin.
> Al momento de presionar enter, podrá visualizar el menú administrador, ver imagen en Administrador.
> Cuenta con dos opciones, las cuales se muestran en la siguiente imagen.
>
> ![menu login!](./imgManuales/login.jpg)
> *Menu principal.*
>

> #### Administrador
>
> Registra los empleados, las imagenes, los clientes registrados, y a los clientes que se atenderán, por medio de archivos csv, por lo que es necesario ingresar la url de la ubicación del archivo, como se muestra en las imagenes.
>
> ![menu administrador!](./imgManuales/cargaEmpleados.jpg)
> *Opción cargar empleados.*
>
> ![menu administrador!](./imgManuales/cargaImagenes.jpg)
> *Opción cargar imagenes.*
>
> ![menu administrador!](./imgManuales/cargarUsuarios.jpg)
> *Opción cargar clientes registrados.*
>
> ![menu administrador!](./imgManuales/actualizarCola.jpg)
> *Opción cargar clientes a la cola.*
>
> La opción 5 permitirá la observación de la imagen de los reportes correspondientes a la estructura utilizada para almacenar cada uno de los datos de las opciones anteriores. ver las siguientes imagenes.
>
> ![menu administrador!](./imgManuales/listadosimple.jpg)
> *Reporte lista simple, correspondiente a empleados.*
>
> ![menu administrador!](./imgManuales/listadoble.jpg)
> *Reporte lista doblemente enlazada, correspondiente a imagenes.*
>
> ![menu administrador!](./imgManuales/listadoCircularSimple.jpg)
> *Reporte lista circular enlazada, correspondiente a clientes registrados.*
>
> ![menu administrador!](./imgManuales/cola.jpg)
> *Reporte cola, correspondiente a clientes a actualizar.*
>

> #### Menu Empleados
> Maneja la visualización de las imagenes, así como el control del pedido de cada uno de los clientes que se encuentran en la cola, y si un cliente no se encuentra registrado, permite agregarlo como cliente nuevo, mostrando su respectivo id, y agregandolo a la lista de clientes registrados. Cuando se inicia sesión automaticamente se mostrará el usuario que realizó el login.
>
> ![menu empleados!](./imgManuales/opcionVerimagenes.jpg)
> *Opción para ver imagenes previas.*
>
> ![menu empleados!](./imgManuales/opcionPedido.jpg)
> *Opción para realizar los pedidos.*
>
> Si el usuario elige la opción número 1, entonces se muestra la lista de imagenes cargadas en el sistema, y podrá elegir cual de estas opciones desea visualizar, y cuando lo seleccione, automaticamente se generará la imagen, ver la siguiente imagen.
>
> ![menu empleados!](./imgManuales/letra_R.jpg)
> *Visualización de la imagen seleccionada.*
>
> Para la opción número 2, correspondiente a realizar pedido, se muestra nuevamente la lista de imagenes y cada vez que un cliente almacenado en la cola elige una opción automaticamente este cliente se elimina de la cola y se almacena en la pila, por lo que se debe seleccionar la imagen para cada cliente que continue en la cola, hasta que esta finalice, y luego se muestra el reporte de la pila con los datos del cliente y la imagen que seleccionó, además taqmbién se genera un archivo JSON con los datos de la pila, ver la siguientes imagenes correspondientes a los reportes.
>
> ![menu empleados!](./imgManuales/pila.jpg)
> *Visualización del reporte con los datos de los clientes con la imagen elegida.*
>
> ![menu empleados!](./imgManuales/reporteJson.jpg)
> *Visualización del reporte JSON con los datos de los clientes con la imagen elegida.*
>

## **MANUAL TECNICO**

EDD Creative maneja la información por medio de carga masiva de archivos de tipo csv, los cuales contienen la información necesaria para la interacción con los empleados. 
Para la creación del software se utiliza el lenguaje de programación go.

> ### Metodo para el menu principal
>
> Para comenzar con la estructura del software se realiza el menú principal tomando en cuenta el inicio de sesión, por lo que se utiliza un bucle para permitir elegir diversas opciones, además  se muestra en pantalla las opciones, y con Scanln se ingresa la opción tecleada por el usuario, de tal manera que accede a la posision de memoria para asignar el dato leído desde consola, luego con el switch evalúa la opcion elegida por el usuario, si es la opcion 1 accede al metodo sesion.
>
```go
func menuPrincipal() {
    var opcion int
    for opcion != 2 {
	    fmt.Println(`
--------- Login ---------
1. Iniciar Sesion
2. Salir del Sistema
-------------------------
Seleccione una opción:`)

		fmt.Scanln(&opcion)

	    switch opcion {
		    case 1:
			    sesion()
		}
	}
}
```

> ### Metodo para sesion
>
> Se crean las variables para el usuario y para la contraseña, se muestra en pantalla los mensajes para ingresarlos, luego se compara con un if si el usuario y la contraseña son los predeterminados para el administrador, si esto es correcto, entonces se abre el menu de administrador, de lo contrario se abre el menu de empleados, por lo tanto se debe validar si el usuario y la contraseña existen en el sistema, por lo cual se envian los datos como parametro hacia la funcion validar para así verificar la existencia de los mismos, por lo que si los datos existen, la función devuelve true, y si no existen en la lista devuelve false, y si devuelve false se muestra un mensaje indicando que el usuario no existe o que se ingresaron mal los datos del mismo. 
>
```go
func sesion() {
	var usuario string
	var password string
	fmt.Println("\nIngrese Usuario: ")
	fmt.Scanln(&usuario)
	fmt.Println("Password: ")
	fmt.Scanln(&password)

	if usuario == "ADMIN_201901103" && password == "Admin" {
		fmt.Println("Bienvenido a admin")
		menuAdministrador()
	} else {

		validandoExistencia := listaSimple.Validar(usuario, password)

		if validandoExistencia == true {
			menuEmpleado(usuario)
		} else {
			fmt.Println("El usuario no existe o ingresó mal el usuario.")
		}
	}
}
```

> ### Metodo para menu administrador
>
> Se crea la variable para seleccionar la opcion, luego se crea un bucle para mantenerse dentro de la sesión, y se muestran las opciones imprimiendolas en consola, luego con la funcion Scanln se recibe la opción para así poder validar a traves del switch a que opción corresponde, por lo que se accede a diferentes metodos los cuales se explicarán más adelante. 
>
```go
func menuAdministrador() {
	var opcion int
	for opcion != 6 {
		fmt.Println(`
--------- Dashboard Administrador 201901103 ---------
1. Cargar Empleados
2. Cargar Imagenes
3. Cargar Usuarios
4. Actualizar Cola
5. Reportes Estructuras
6. Cerrar Sesion
-----------------------------------------------------
Seleccione una opción:`)

		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			cargarEmpleados()
		case 2:
			cargarImagenes()
		case 3:
			cargarClientes()
		case 4:
			cargarActualizarCola()
		case 5:
			listaSimple.ReporteSimple()
			listaDoble.ReporteDoble()
			listaCircular.ReporteCircular()
			clientesCola.ReporteCola()
		}
	}
}
```

> ### Metodo para cargar empleados
>
> Se crea la variable para la ruta que ingresa el usuario, de esa manera se lee el archivo correctamente, por lo que se utiliza la funcion Open, la cual recibe como parametro la ruta del archivo, y si existe algún error con el archivo, entonces se muestra un mensaje indicando que ha ocurrido un error, y de lo contrario con defer file close se cierra el archivo, y con la funcion transform NewReader se garantiza que el archivo lea los datos independientemente de los tipos de caracteres que incluya el mismo, por lo que esta variable se envía a NewReader del csv para así leer todo el archivo, luego se indica con Comma cual será el caracter que separa los datos, despues con RedAll se leen todas las lineas del archivo, luego si ocurre algun error al leer las lineas se muestra un mensaje en pantalla con el error que ha ocurrido, si no existe ningun error, entonces continua con el bucle, el cual se encarga de recorrer cada dato de cada linea y así poder validar con el if si la linea es la cabecera, entonces lo omite, de lo contrario accede a cada valor, y se envia como parametro al metodo Insertar de la lista simple, indicando con TrimSpace que los datos no tengan espacios extras.
>
```go
func cargarEmpleados() {
	var ruta string
	fmt.Println("Ingrese la ruta del archivo: ")
	fmt.Scanln(&ruta)

	// Abre el archivo CSV
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Crea un lector con transformador UTF-8
	utf8Reader := transform.NewReader(file, unicode.UTF8.NewDecoder())

	// Crea un nuevo lector CSV
	reader := csv.NewReader(utf8Reader)
	reader.Comma = ','

	// Lee todas las líneas del archivo
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	// Itera sobre las líneas y muestra los datos
	for _, line := range lines {
		if line[0] != "id" {
			//fmt.Println(line[0], " ", line[1], " ", line[2], " ", line[3])
			listaSimple.Insertar(strings.TrimSpace(line[0]), strings.TrimSpace(line[1]), strings.TrimSpace(line[2]), strings.TrimSpace(line[3]))
		}
	}
	//listaSimple.Mostrar()
}
```

> ### Metodo para cargar imagenes
>
> Se crea una variable que obtendrá la ruta del archivo de imagenes csv, el cual será abierto de la misma manera como se explica en el metodo anterior, con la variacion que ahora los datos se insertan en la lista doble, ademas se crea una bandera de encabezado, para omitir el mismo, cuando se encuentre en el archivo.
>
```go
func cargarImagenes() {
	var ruta string
	fmt.Println("Ingrese la ruta del archivo: ")
	fmt.Scanln(&ruta)

	// Abre el archivo CSV
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Crea un lector con transformador UTF-8
	utf8Reader := transform.NewReader(file, unicode.UTF8.NewDecoder())

	// Crea un nuevo lector CSV
	reader := csv.NewReader(utf8Reader)
	reader.Comma = ','
	encabezado := true

	for {
		lines, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error al leer la linea del archivo")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		listaDoble.Insertar(strings.TrimSpace(lines[0]), strings.TrimSpace(lines[1]))
	}
}
```

> ### Metodo para cargar clientes
>
> Se crea una variable que obtendrá la ruta del archivo de imagenes csv, el cual será abierto de la misma manera como se explica en el metodo de cargar empleados, con la variacion que ahora los datos se insertan en la lista circular, ademas se crea una bandera de encabezado, para omitir el mismo, cuando se encuentre en el archivo.
>
```go
func cargarClientes() {
	var ruta string
	fmt.Println("Ingrese la ruta del archivo: ")
	fmt.Scanln(&ruta)

	// Abre el archivo CSV
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Crea un lector con transformador UTF-8
	utf8Reader := transform.NewReader(file, unicode.UTF8.NewDecoder())

	// Crea un nuevo lector CSV
	reader := csv.NewReader(utf8Reader)
	reader.Comma = ','
	encabezado := true

	for {
		lines, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error al leer la linea del archivo")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		listaCircular.Insertar(strings.TrimSpace(lines[0]), strings.TrimSpace(lines[1]))
	}
}
```

> ### Metodo para cargar clientes en cola
>
> Se crea una variable que obtendrá la ruta del archivo de imagenes csv, el cual será abierto de la misma manera como se explica en el metodo de cargar empleados, con la variacion que ahora los datos se insertan en la cola, ademas se crea una bandera de encabezado, para omitir el mismo, cuando se encuentre en el archivo.
>
```go
func cargarActualizarCola() {
	var ruta string
	fmt.Println("Ingrese la ruta del archivo: ")
	fmt.Scanln(&ruta)

	// Abre el archivo CSV
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Crea un lector con transformador UTF-8
	utf8Reader := transform.NewReader(file, unicode.UTF8.NewDecoder())

	// Crea un nuevo lector CSV
	reader := csv.NewReader(utf8Reader)
	reader.Comma = ','
	encabezado := true

	for {
		lines, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error al leer la linea del archivo")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		clientesCola.Encolar(strings.TrimSpace(lines[0]), strings.TrimSpace(lines[1]))
	}
}
```

> ### Metodo para generar reportes
>
> Se crea
>
```go
```

> ### Metodo para el menu de empleados
>
> Se crea la variable opcion para manejar la entrada por consola escrita por el usuario, por lo tanto con el bucle for se indica que es diferente de la opcion cerrar sesión, para así poder regresar al menú principal ya que finaliza la ejecución del bucle, además con el switch verifica a que opción corresponde, las opciones se verán más adelante.
>
```go
func menuEmpleado(usuario string) {
	var opcion int
	for opcion != 4 {
		fmt.Printf(`
--------- EDD Creative %s ---------
1. Ver Imagenes Cargadas
2. Realizar Pedido
3. Capas
4. Cerrar Sesion
-----------------------------------------------------
Seleccione una opción:`, usuario)

		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			nameImagen := visualizarImagenes()
			fmt.Println("La imagen elegida fue: ", nameImagen, "\nMostrando visualizacion previa")
			previaVisualizacion(nameImagen)
		case 2:
			realizarPedidos(usuario)
			pedidosPila.ReportePila()
			pedidosPila.ReporteJson()
		case 3:
			nameImagen := visualizarImagenes()
			fmt.Println("La imagen elegida fue: ", nameImagen, "\nMostrando visualizacion previa")
			realizarCapa(nameImagen)
		}

	}
}
```

> ### Función para visualizar las imagenes
>
> Se crea una opción para validar la entrada del usuario, luego se manda a llamar la lista de datos correspondiente a las imagenes, las cuales estan almacenadas en la lista doble, además despues de haber obtenido la opcion, esta se envia como parametro a traves de la funcion BuscarImagen, y devuelve el nombre de la imagen y luego retorna este nombre para almacenarla en la variable desde donde se realiza la invocación a la función.
>
```go
func visualizarImagenes() string {
	var opcion int
	fmt.Println("\n###################Listado de Imagenes###################")
	listaDoble.ListarDatos()
	fmt.Println("\n Seleccione una opción:")
	fmt.Scanln(&opcion)
	nameImagen := listaDoble.BuscarImagen(strconv.Itoa(opcion))
	return nameImagen
	//Falta la opcion de visualizar la imagen
}
```

> ### Función para la primera opcion correspondiente a visualización previa
>
> Se crea la variable de matriz, la cual inicializará los valores de la raiz hacia el nodo de la matriz, con los datos de posiciones en -1 debido a que todos los datos dentro de la matriz pueden iniciar con el nodo 0 e ir aumentando conforme sea necesario, y por esa razón tampoco se le indica un color en especifico, además se envía de manera inicial por medio de parametro la ruta correspondiente a la ubicación de la carpeta con las configuraciones para las imagenes, esta es csv más el nombre de la imagen recibida por parametro y el inicial.csv para la lectura de las capas que tendrá el archivo, y nuevamente se agrega el nombre de la imagen la cual será para la extensión del archivo css y html, por lo que se manda a llamar el metodo GenerarImagen, luego simplemente se inicializa la matriz, para volver a generar otra imagen de ser necesario.
>
```go
func previaVisualizacion(nameImagen string) {
	var matrizImages = &estructura.Matriz{Raiz: &estructura.NodoMatriz{PosicionX: -1, PosicionY: -1, Color: "RAIZ"}}
	matrizImages.LeerInicial("csv/"+nameImagen+"/inicial.csv", nameImagen)
	matrizImages.GenerarImagen(nameImagen)
	matrizImages = &estructura.Matriz{Raiz: nil}
}
```

> ### Metodo para realizar pedidos
>
> Se crea un bucle infinito que sirve para obtener los datos actuales de la cola, en este caso el id y nombre del cliente, así como el tamaño de la misma, luego con un if se valida si la longitud de la cola es diferente de 0, esto para indicar que cuando sea igual a 0 se finaliza el bucle, luego se muestra un mensaje indicando cual es el cliente que se está atendiendo, se valida si el cliente no está registrado en la lista circular, por lo que se crea otro bucle dentro de la validación cuando sea igual a X, este bucle servirá para que se repita la generación del id en caso de que este exista, de esta forma aseguramos que solo existan id unicos en el sistema, luego se valida cuando el usuario no existe, entonces se muestran las imagenes disponibles en el sistema, por lo que el usuario puede elegir una imagen y luego se ingresa el cliente nuevo a la lista circular, además se agrega el id del cliente, el id del empleado y la imagen elegida a la pila, luego se indica por medio de un mensaje cual es el id para determinado cliente, y se saca de la cola al cliente atendido y se utiliza un break para romper el bucle, luego en el else correspondiente a la validación si es diferente de X, se vuelve a validar si el cliente se encuentra en el sistema, y se visualizan las imagenes, se agrega el usuario a la pila, y se elimina el usuario de la cola, luego si no existe el cliente, simplemente se agrega a la lista circular y se vuelven a realizar las asignaciones.
>
```go
func realizarPedidos(usuario string) {
	for {
		idcolaClientes := clientesCola.ObtenerClienteId()
		nameColaClientes := clientesCola.ObtenerClienteName()
		longi := clientesCola.ObtenerLongitud()
		if longi != 0 {
			fmt.Println("\nAtendiendo al cliente con id: ", idcolaClientes, " y nombre: ", nameColaClientes)

			if strings.ToUpper(idcolaClientes) == "X" {
				// CUANDO ES IGUAL A X VALIDAR UN ID RANDOM Y
				for {
					valor := (rand.Intn(10000)) + 1000

					existe := listaCircular.ValidarRepetidos(strconv.Itoa(valor))
					if existe == true {
						//repetir el aleatorio y no guardar nada
					} else {
						// guardar el aleatorio como nuevo id y agregarlo a la lista circular junto al nombre del cliente
						nombreImagenElegida := visualizarImagenes()
						//Sino existe en la lista circular agregar al cliente en la lista circular
						listaCircular.Insertar(strconv.Itoa(valor), nameColaClientes)
						pedidosPila.Push(strconv.Itoa(valor), usuario, nombreImagenElegida)
						//agregar el id del cliente, id del empleado, y nombre de la imagen elegida
						fmt.Println("\nEl nuevo id: ", strconv.Itoa(valor), "corresponde al cliente: ", nameColaClientes)
						clientesCola.Descolar()
						break
					}
				}

			} else {
				existe := listaCircular.ValidarRepetidos(strings.TrimSpace(idcolaClientes))
				if existe == true {
					// si el cliente existe en la lista circular de clientes
					nombreImagenElegida := visualizarImagenes()
					pedidosPila.Push(idcolaClientes, usuario, nombreImagenElegida)
					//agregar el id del cliente, id del empleado, y nombre de la imagen elegida
					clientesCola.Descolar()
				} else {
					nombreImagenElegida := visualizarImagenes()
					//Sino existe en la lista circular agregar al cliente en la lista circular
					listaCircular.Insertar(idcolaClientes, nameColaClientes)
					pedidosPila.Push(idcolaClientes, usuario, nombreImagenElegida)
					//agregar el id del cliente, id del empleado, y nombre de la imagen elegida
					clientesCola.Descolar()
				}

			}
			fmt.Println("\nFinaliza atención a cliente actual y quedan:", strconv.Itoa(longi-1))
		} else {
			break
		}
	}
}
```

> ### Metodo para generar las capas
>
> Se crea nuevamente la inicialización de la matriz, haciendo la creación del nodo raiz con las posiciones en -1 y color como raiz, luego se crea una lista simple para almacenar las capas, se envia por parametros al meotodo leerInicial1 la ruta del archivo con el nombre de la imagen, y la lista simple de capas, esto solamente sirve para obtener cada capa y luego poder seleccionar solamente una, entonces se vuelve a inicializar la matriz enviandole nulo a la raiz. Despues se crea una variable de opcion la cual servirá directamente para poder elegir la capa a visualizar, por lo que se listan las capas que corresponden a la imagen elegida, luego se busca el nombre de la capa en la lista simple, y se alamcena en la variable nameCapa, se inicia nuevamente la matriz, y se manda a llamar al metodo leer inicial y capa elegida, esto para enviarle el nombre de la imagen, la ruta de la imagen, y el nombre de la capa, y así generar el archivo correspondiente al reporte de la matriz por capas, luego se inicializa nuevamente la matriz.
>
```go
func realizarCapa(nameImagen string) {
	var matrizImages1 = &estructura.Matriz{Raiz: &estructura.NodoMatriz{PosicionX: -1, PosicionY: -1, Color: "RAIZ"}}
	var listaCapasMatriz = estructura.NewListaSimpleCapa()
	matrizImages1.LeerInicial1("csv/"+nameImagen+"/inicial.csv", nameImagen, listaCapasMatriz)
	matrizImages1 = &estructura.Matriz{Raiz: nil}

	var opcion int
	fmt.Println("\n=================Listado de Capas=================")
	listaCapasMatriz.ListarDatosCapa()
	fmt.Println("\n Seleccione una opción:")
	fmt.Scanln(&opcion)
	nameCapa := listaCapasMatriz.BuscarCapa(strconv.Itoa(opcion))
	matrizImages1 = &estructura.Matriz{Raiz: &estructura.NodoMatriz{PosicionX: -1, PosicionY: -1, Color: "RAIZ"}}
	matrizImages1.LeerInicialYCapaElegida("csv/"+nameImagen+"/inicial.csv", nameImagen, nameCapa)
	matrizImages1 = &estructura.Matriz{Raiz: nil}
}
```

> ### Metodo para generar reportes
>
> Se crea
>
```go
```