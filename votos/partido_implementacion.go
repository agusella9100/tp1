package votos

import (
	"bufio"
	"fmt"
	"os"
	TDACola "tdas/cola"
	TDALista "tdas/lista"
)

const ruta = "archivo.txt"

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
	
	//leo el archivo de los partidos, guardo su nombre y los candidatos en una cola, pila o lista
	//cada partido tiene su nombre y 3 candidatos
	return nil
}

func CrearVotosEnBlanco() Partido {
	return new(partidoEnBlanco)
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {

}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {

	return partido.postulantes[tipo]
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {

}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	return ""
}
