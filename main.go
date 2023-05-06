package main

import (
	"bufio"
	"fmt"
	"os"
	TDAOrdenYBusqueda "rerepolez/ordenamiento"
	TDAVotos "rerepolez/votos"
	"strings"
	TDACola "tdas/cola"
	//"strings"
	//TDACola "tdas/cola"
	"strconv"
)

func main() {

	padrones, errdni := os.Open(os.Args[1])
	partidos, errPartidos := os.Open(os.Args[2])

	if errdni != nil {
		fmt.Printf("Archivo %s no encontrado\n", padrones)
		return
	}
	defer padrones.Close()

	d := bufio.NewScanner(padrones)
	nroVotantes := 0
	arrayVotantes := []TDAVotos.Votante{}

	for d.Scan() {
		dni, errAtoi := strconv.Atoi(d.Text())
		votante := TDAVotos.CrearVotante(dni)

		arrayVotantes = append(arrayVotantes, votante)

		nroVotantes++
		if errAtoi != nil {
			fmt.Println(errAtoi)
		}
	}

	errdni = d.Err()
	if errdni != nil {
		fmt.Println(errdni)
	}
	TDAOrdenYBusqueda.OrdenarPadrones(arrayVotantes)
	/*for i := 0; i < nroVotantes; i++ {
		fmt.Println(arrayVotantes[i].LeerDNI())
	}*/

	p := bufio.NewScanner(partidos)

	if errPartidos != nil {
		fmt.Printf("Archivo %s no encontrado\n", partidos)
		return
	}

	defer partidos.Close()

	//creo array de partidos.
	//puedo hacerlo una lista. lo dejo comentado esta opción
	//La ventaja que veo de hacer listas es que guardar los partidos es O(1)
	//ya que la primitiva insertar es de orden cte mientras que append ni idea, creo es O(n).

	//listaPartidos := TDALista.CrearListaEnlazada[TDAVotos.Partido]()
	//listaPartidos.InsertarPrimero(TDAVotos.CrearVotosEnBlanco())
	arrPartidos := []TDAVotos.Partido{}
	arrPartidos = append(arrPartidos, TDAVotos.CrearVotosEnBlanco())

	for p.Scan() {
		linea := p.Text()
		nombresPartidos := strings.Split(linea, ",")
		nombre := nombresPartidos[0]
		candidatos := [3]string{nombresPartidos[1], nombresPartidos[2], nombresPartidos[3]}
		partidoCreado := TDAVotos.CrearPartido(nombre, candidatos)
		//listaPartidos.InsertarUltimo(partidoCreado)
		arrPartidos = append(arrPartidos, partidoCreado)
	}
	errPartidos = p.Err()
	if errPartidos != nil {
		fmt.Println(errPartidos)
	}

	scanner := bufio.NewScanner(os.Stdin)
	colaIngresantes := TDACola.CrearColaEnlazada[TDAVotos.Votante]()

	//For que termina cuando finaliza entrada estandar

	//hace la llamada en cada linea, para analizar los comandos.
	//Habria que definir su ComandosLeidos devuelve algo(un error) o es void
	for scanner.Scan() {
		linea := scanner.Text()
		TDAOrdenYBusqueda.ComandosLeidos(linea, colaIngresantes, arrayVotantes, arrPartidos)
	}

}
