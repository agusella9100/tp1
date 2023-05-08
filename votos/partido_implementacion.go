package votos

import (
	"strconv"
)

type partidoImplementacion struct {
	nombre      string
	postulantes [CANT_VOTACION]string
	contadores  [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	contadores [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(partidoImplementacion)
	partido.nombre = nombre

	partido.postulantes[PRESIDENTE] = candidatos[PRESIDENTE]
	partido.postulantes[GOBERNADOR] = candidatos[GOBERNADOR]
	partido.postulantes[INTENDENTE] = candidatos[INTENDENTE]

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

	if votos != "1" {
		devolucion = devolucion + " votos"
	} else {
		devolucion = devolucion + " voto"
	}
	return devolucion
}

// Pienso que se repite un poco el codigo en la implementacion de las primitivas para cada partido
// por ahi conviene crear una funcion que llamamos en las primitivas y que devuelva el mensaje
func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.contadores[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	votos := strconv.Itoa(blanco.contadores[tipo])
	devolucion := "Votos en blanco: " + votos
	uno := strconv.Itoa(1)

	if votos != uno {
		devolucion = devolucion + " votos"
	} else {
		devolucion = devolucion + " voto"
	}
	return devolucion

}
