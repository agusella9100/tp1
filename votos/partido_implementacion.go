package votos

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	TDACola "tdas/cola"
	TDALista "tdas/lista"
)

const ruta = "archivo.txt"

//Esto supongo que ira en un archivo go en el que implementamos todo el sistema de votos llamando a las primitivas y todo eso
// como es para leer un archivo segun apuntes de la catedra
func leerArchivo() {
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

// esto eera para probar el import, despues lo borramos cuando codiemos cosas donde usemos las listas
func list() {
	l := TDALista.CrearListaEnlazada[int]()
	l.InsertarPrimero(5)
	c := TDACola.CrearColaEnlazada[int]()
	c.Encolar(5)
}

type partidoImplementacion struct {
	nombre      string
	postulantes [CANT_VOTACION]string
	contadores  [CANT_VOTACION]int
	numlista int
}

type partidoEnBlanco struct {
	contadores [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(partidoImplementacion)
	partido.nombre = nombre

	//leo el archivo de los partidos, guardo su nombre y los candidatos en una cola, pila o lista
	//cada partido tiene su nombre y 3 candidatos
	return partido
}

func CrearVotosEnBlanco() Partido {
	return new(partidoEnBlanco)
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.contadores[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	votos := strconv.Itoa(partido.contadores[tipo])
	devolucion := partido.nombre + " - " + partido.postulantes[tipo] + ": " + votos

	if votos != 1 {
		devolucion = devolucion + " votos"
	}else {
		devolucion = devolucion + " voto"
	}
	return devolucion
}

//Pienso que se repite un poco el codigo en la implementacion de las primitivas para cada partido
//por ahi conviene crear una funcion que llamamos en las primitivas y que devuelva el mensaje
func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.contadores[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	votos := strconv.Itoa(blanco.contadores[tipo])
	devolucion := "Votos en blanco: " + votos

	if votos != 1 {
		devolucion = devolucion + " votos"
	}else {
		devolucion = devolucion + " voto"
	}
	return devolucion

}
