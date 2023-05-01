package votos

import (
	"bufio"
	"fmt"
	"os"
	TDACola "tdas/cola"
	TDALista "tdas/lista"
)

const ruta = "archivo.txt"

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
	postulantes [3]string
	contadores  [3]int
}

type partidoEnBlanco struct {
	contadores [3]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(partidoImplementacion)
	partido.nombre = nombre
	partido.postulantes = candidatos
	return partido
}

func CrearVotosEnBlanco() Partido {
	return new(partidoEnBlanco)
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.contadores[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	return
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.contadores[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	return ""
}
