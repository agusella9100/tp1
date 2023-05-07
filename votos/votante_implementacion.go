package votos

type votanteImplementacion struct {
    dni    int
    yaVoto bool
    votosRealizados pilaDinamica[Voto]
}

func CrearVotante(dni int) Votante {
    nuevoVoto := new(votanteImplementacion)
    nuevoVoto.votosRealizados = CrearPilaDinamica[Voto]()
    return nuevoVoto
}

func (votante votanteImplementacion) LeerDNI() int {
    return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
    //Si ya voto devuelvo el error
    if votante.yaVoto {
        votante.votosRealizados.Apilar(new(Voto))
        return ErrorVotanteFraudulento{Dni: votante.dni}
    }

    //Aca creo un nuevo voto, copio el anterior y modifico lo que corresponde
    voto := votante.votosRealizados.VerTope()
    voto.VotoPorTipo[tipo] = alternativa
    if alternativa == 0 {
        voto.Impugnado = true
    }
    votante.votosRealizados.Apilar(voto)

    return nil

}

func (votante *votanteImplementacion) Deshacer() error {
    //Si no hay votos, imprime el error
    if votante.votosRealizados.EstaVacia() {
        return ErrorNoHayVotosAnteriores{}
    }
    //Si ya voto
    if votante.yaVoto {
        votante.votosRealizados.Apilar(new(Voto))
        return ErrorVotanteFraudulento{Dni: votante.dni}
    }

    //Elimino la version del ultimo voto
    votante.votosRealizados.Desapilar()


    return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
    if votante.yaVoto {
        votante.votosRealizados.Apilar(new(Voto))
        return Voto{}, ErrorVotanteFraudulento{Dni: votante.dni}
    }

    votante.yaVoto = true

    return votante.votosRealizados.VerTope(), nil
}
