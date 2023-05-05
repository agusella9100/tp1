package main

import (
	"bufio"
	"fmt"
	"os"
	TDAOrdenYBusqueda "rerepolez/ordenamiento"
	TDAVotos "rerepolez/votos"
	//TDAPila "tdas/pila"
	"strconv"
)

func main() {
	//todo averiguar como recibir los archivos como parametros asi se leen despues
	padrones, errdni := os.Open("tests/02_padron")

	if errdni != nil {
		fmt.Printf("Archivo %s no encontrado\n", padrones)
		return
	}
	defer padrones.Close()

	d := bufio.NewScanner(padrones)
	nroVotantes := 0
	arrayVotantes := []TDAVotos.Votante{}
	fmt.Println("elementos votantes ", len(arrayVotantes))

	for d.Scan() {
		dni, errAtoi := strconv.Atoi(d.Text())
		votante := TDAVotos.CrearVotante(dni)
		fmt.Printf("Leí: %v\n", dni)
		//fmt.Printf("De la fila: %v \n", nroVotantes)
		arrayVotantes = append(arrayVotantes, votante)
		//pilaVotantes.Apilar(votante)
		nroVotantes++
		if errAtoi != nil {
			fmt.Println(errAtoi)
		}
	}

	errdni = d.Err()
	if errdni != nil {
		fmt.Println(errdni)
	}
	//listo ya funciona el cargar los dnis y ordenarlos:
	fmt.Println("elementos votantes ", len(arrayVotantes))
	TDAOrdenYBusqueda.OrdenarPadrones(arrayVotantes)
	for i := 0; i < nroVotantes; i++ {
		fmt.Println(arrayVotantes[i].LeerDNI())
	}

	//todo leer documento partidos y guardar la información como corresponde.
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
