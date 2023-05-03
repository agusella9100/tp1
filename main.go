package tp1

import (
	"bufio"
	"fmt"
	"os"
)

const ruta = "archivo.txt"

func main() {

	//Esto supongo que ira en un archivo go en el que implementamos todo el sistema de votos llamando a las primitivas y todo eso
	// como es para leer un archivo segun apuntes de la catedra
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Printf("Archivo %s no encontrado", ruta)
		return
	}
	defer archivo.Close()

	s := bufio.NewScanner(archivo)
	for s.Scan() {
		fmt.Printf("Leí: %s\n", s.Text())
	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
	}
}
