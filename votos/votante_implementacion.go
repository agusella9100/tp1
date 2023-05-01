package votos

type votanteImplementacion struct {
	dni    int
	yaVoto bool
	suVoto Voto
}

func CrearVotante(dni int) Votante {
	nuevoVoto := new(votanteImplementacion)
	nuevoVoto.dni = dni
	return nuevoVoto
}

func (votante votanteImplementacion) LeerDNI() int {
	//supongo valida que el dni
	return 0
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	return Voto{}, nil
}
