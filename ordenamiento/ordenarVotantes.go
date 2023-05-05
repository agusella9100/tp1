package ordenamiento

import (
	TDAVotante "rerepolez/votos"
)

// es un quickSort: recive un array de votantes (primitivas leerdni votar-deshacer-finvoto) un votante tiene variables (dni-yavoto-votosrealizados)
func OrdenarPadrones(arr []TDAVotante.Votante) []TDAVotante.Votante {
	//caso base
	if len(arr) < 2 {
		return arr
	}
	//caso general
	//medio := len(arr) / 2
	pivote := 0
	ultimoMenor := 0
	//acordarse de al swapear, pasar la dirección
	//swap(&arr[pivote], &arr[medio]) // esto dicen es para optimizar

	for i := pivote + 1; i < len(arr); i++ {
		if arr[i].LeerDNI() < arr[pivote].LeerDNI() {
			ultimoMenor++
			swap(&arr[i], &arr[ultimoMenor]) //aca voy acomodando todos los menores al pivote en fila
		}
	}
	//coloco al elemento del pivote al final de la fila de los menores al pivote y al ultimomenor en la primera posición
	swap(&arr[ultimoMenor], &arr[pivote])

	//hago el llamado a la recursividad para que se ordenen las mitades menores al elemento del pivote y mayores al mismo
	//acordarse que los slices no incluyen al final
	OrdenarPadrones(arr[0 : ultimoMenor+1]) //hago el llamado para la mitad de menores hasta donde está el pivote
	OrdenarPadrones(arr[ultimoMenor+1:])    //hago el llamado para la mitad superior al pivote.
	return arr
}

func swap(a *TDAVotante.Votante, b *TDAVotante.Votante) {
	//capaz sale si hago que se pase como parametro la dirección

	*a, *b = *b, *a

	/*
		votanteAux := a
		a = b
		b = votanteAux */
	// con esto no es que cambio exactamente, sino que sobreescribo lo que tenia en cada posición
}
