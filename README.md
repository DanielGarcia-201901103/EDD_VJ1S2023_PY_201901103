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

```go
func menuPrincipal() {
    var opcion int
    for opcion != 2 {
	    fmt.Println(`
--------- Login ---------
1. Iniciar Sesion
2. Salir del Sistema
-------------------------
Elige una opción:`)

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
		//buscar entre los empleados guardados en la lista simple enlazada y si existe iniciar sesion en menu empleado
		/*Se envia usuario y password por parametro al metodo buscar en la lista simple enlazada y si el usuario y el password
		coinciden entonces devuelve true y inicia sesion, de lo contrario devuelve false y muestra el mensaje*/
		//menuEmpleado()
		//si no coincide entonces mostrar el siguiente mensaje
		//fmt.Println("El usuario no existe")
	}
}
```