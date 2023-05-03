package rerepolez

import (
	"bufio"
	"fmt"
	"os"
)

const ruta = "archivo.txt"

func main() {

	//Esto supongo que ira en un archivo go en el que implementamos todo el sistema de votos llamando a las primitivas y todo eso
	// como es para leer un archivo segun apuntes de la catedra

	s := bufio.Scanner{}
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Printf("Archivo %s no encontrado", ruta)
		return
	}
	defer archivo.Close()

	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}

		err = s.Err()
		if err != nil {
			fmt.Println(err)
		}
		return 0, data, bufio.ErrFinalToken
	}
	a := bufio.NewScanner(os.Stdin)
	a.Split(onComma)
	columna := 0
	for a.Scan() {
		fmt.Printf("Leí: %s\n", a.Text())
		fmt.Printf("De la columna: %v \n", columna)
	}
}
