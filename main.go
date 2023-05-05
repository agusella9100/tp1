package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	padrones, errdni := os.Open("tests/02_padron")

	if errdni != nil {
		fmt.Printf("Archivo %s no encontrado\n", padrones)
		return
	}
	defer padrones.Close()

	d := bufio.NewScanner(padrones)
	fila := 0
	for d.Scan() {
		fmt.Printf("Leí: %s\n", d.Text())
		fmt.Printf("De la fila: %v \n", fila)
		fila++
	}
	errdni = d.Err()
	if errdni != nil {
		fmt.Println(errdni)
	}

	partidos, errPartidos := os.Open("tests/02_partidos")
	p := bufio.NewScanner(partidos)
	p2 := bufio.NewScanner(partidos)

	if errPartidos != nil {
		fmt.Printf("Archivo %s no encontrado\n", partidos)
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
		err = p2.Err()
		if err != nil {
			fmt.Println(err)
		}
		return 0, data, bufio.ErrFinalToken
	}

	p2.Split(onComma)

	columna := 0

	for p2.Scan() {
		fmt.Printf("De la columna: %v \n", columna)
		fmt.Printf("Leí: %s\n", p2.Text())
		columna++
		if columna > 3 {
			columna = 0
		}
	}
	fmt.Println("salto de linea")

	columna = 0
	errPartidos = p.Err()
	if errPartidos != nil {
		fmt.Println(errPartidos)
	}

}
