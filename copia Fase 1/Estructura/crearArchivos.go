package estructura

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func crearArchivo(nombre_Archivo string) {
	var _, err = os.Stat(nombre_Archivo)
	if os.IsNotExist(err) {
		var file, err = os.Create(nombre_Archivo)
		if err != nil {
			return
		}
		defer file.Close()
	}
	fmt.Println("Archivo generado exitosamente")
}

// Escribir archivo
func escribirArchivo(contenido string, nombrer_Archivo string) {
	var file, err = os.OpenFile(nombrer_Archivo, os.O_RDWR, 0644) //os.RDWR cambia los permisos
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.WriteString(contenido)
	if err != nil {
		return
	}
	err = file.Sync()
	if err != nil {
		return
	}
	fmt.Println("Archivo guardado exitosamente")
}

func ejecutar(nombre_imagen string, archivo string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tjpg", archivo).Output()
	mode := 0777
	_ = ioutil.WriteFile(nombre_imagen, cmd, os.FileMode(mode))

}
