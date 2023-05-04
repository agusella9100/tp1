package rerepolez

import (
	"bufio"
	"fmt"
	"os"
)

func main(argc int, argv []string) {

	fmt.Printf("recibi %v cosas \n", argc)
	//Esto supongo que ira en un archivo go en el que implementamos todo el sistema de votos llamando a las primitivas y todo eso
	// como es para leer un archivo segun apuntes de la catedra

	d := bufio.Scanner{}
	padrones, errd := os.Open(argv[1])

	if errd != nil {
		fmt.Printf("Archivo %s no encontrado\n", argv[1])
		return
	}
	defer padrones.Close()

	p := bufio.Scanner{}
	partidos, errPartidos := os.Open(argv[2])

	if errPartidos != nil {
		fmt.Printf("Archivo %s no encontrado\n", argv[2])
		return
	}

	defer partidos.Close()

	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		err = p.Err()
		if err != nil {
			fmt.Println(err)
		}
		return 0, data, bufio.ErrFinalToken
	}

	p.Split(onComma)

	columna := 0
	for p.Scan() {
		fmt.Printf("Leí: %s\n", p.Text())
		fmt.Printf("De la columna: %v \n", columna)
		columna++
	}
	fila := 0
	for d.Scan() {
		fmt.Printf("Leí: %s\n", d.Text())
		fmt.Printf("De la fila: %v \n", fila)
		fila++
	}

}
