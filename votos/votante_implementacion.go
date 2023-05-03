package votos

import (
	"rerepolez/errores"
	TDAPila "tdas/pila"
)

type votanteImplementacion struct {
	dni             int
	yaVoto          bool
	votosRealizados TDAPila.Pila[Voto]
}

func CrearVotante(dni int) Votante {
	nuevoVoto := new(votanteImplementacion)
	nuevoVoto.votosRealizados = TDAPila.CrearPilaDinamica[Voto]()
	nuevoVoto.dni = dni
	return nuevoVoto
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	//Si ya voto devuelvo el error
	if votante.yaVoto {
		//fraude := Errores.ErrorVotanteFraudulento{Dni: votante.dni}
		/*errorFraude := new(errores.ErrorVotanteFraudulento)
		errorFraude.Dni = votante.dni*/
		return errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}
	//Aca creo un nuevo voto, copio el anterior y modifico lo que corresponde
	voto := votante.votosRealizados.VerTope()
	voto.VotoPorTipo[tipo] = alternativa
	votante.votosRealizados.Apilar(voto)

	return nil

}

func (votante *votanteImplementacion) Deshacer() error {
	//Si no hay votos, imprime el error
	if votante.votosRealizados.EstaVacia() {

		return errores.ErrorNoHayVotosAnteriores{}
	}
	//Si ya voto
	if votante.yaVoto {
		fraude := errores.ErrorVotanteFraudulento{Dni: votante.dni}
		fraude.Error() // capaz no hay que hacer esto sino que en el "main" si recivimos un error != nil habria que aplicar el metodo Error()
		return fraude
	}

	//Elimino la version del ultimo voto
	votante.votosRealizados.Desapilar()

	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if votante.yaVoto {
		fraude := errores.ErrorVotanteFraudulento{Dni: votante.dni}
		fraude.Error()
		return Voto{}, fraude
	}

	return votante.votosRealizados.VerTope(), nil
}
