package ordenamiento

import (
	"fmt"
	"rerepolez/errores"
	TDAVotos "rerepolez/votos"
	"strconv"
	"strings"
	TDACola "tdas/cola"
)

// es un quickSort: recive un array de votantes (primitivas leerdni votar-deshacer-finvoto) un votante tiene variables (dni-yavoto-votosrealizados)
func OrdenarPadrones(arr []TDAVotos.Votante) []TDAVotos.Votante {
	//caso base
	if len(arr) < 2 {
		return arr
	}

	pivote := 0
	ultimoMenor := 0

	for i := pivote + 1; i < len(arr); i++ {
		if arr[i].LeerDNI() < arr[pivote].LeerDNI() {
			ultimoMenor++
			swap(&arr[i], &arr[ultimoMenor]) //aca voy acomodando todos los menores al pivote en fila
		}
	}
	//coloco al elemento del pivote al final de la fila de los menores al pivote y al ultimo menor en la primera posición
	swap(&arr[ultimoMenor], &arr[pivote])

	//hago el llamado a la recursividad para que se ordenen las mitades menores al elemento del pivote y mayores al mismo
	//acordarse que los slices no incluyen al final
	OrdenarPadrones(arr[0 : ultimoMenor+1]) //hago el llamado para la mitad de menores hasta donde está el pivote
	OrdenarPadrones(arr[ultimoMenor+1:])    //hago el llamado para la mitad superior al pivote.
	return arr
}

func swap(a *TDAVotos.Votante, b *TDAVotos.Votante) {
	*a, *b = *b, *a
}

// Una funcion que me dice si el dni esta dentro del padron o no y me devuelve la posicion dentro del array
func esVotanteValido(dni int, votantes []TDAVotos.Votante) (bool, int) {
	inicio := 0
	fin := len(votantes) - 1

	for inicio <= fin {
		medio := (inicio + fin) / 2
		if votantes[medio].LeerDNI() == dni {
			return true, medio
		} else if votantes[medio].LeerDNI() < dni {
			inicio = medio + 1
		} else {
			fin = medio - 1
		}
	}
	return false, -1
}

// lectura de comandos

// recibe la linea leida, una cola donde se guardan los ingresantes que van a votar, el array del padron ordenado (todos los dnis) y un array de los partidos, que podemos hacerlo que sea una lista correctamente.
func ComandosLeidos(linea string, cola TDACola.Cola[TDAVotos.Votante], arrayVotantes []TDAVotos.Votante, listaPartidos []TDAVotos.Partido) {
	comandos := strings.Split(linea, " ")

	// discriminación del comando

	if comandos[0] == "ingresar" {
		if len(comandos) != 2 {
			imprimirErrores(new(errores.ErrorParametros))
			//continue
		}
		dni, err := strconv.Atoi(comandos[1])
		if err != nil {
			imprimirErrores(new(errores.DNIError))
			//continue
		}

		if dni < 0 {
			imprimirErrores(new(errores.DNIError))
			//continue
		}

		estaDentroDelPadron, _ := esVotanteValido(dni, arrayVotantes)

		if !estaDentroDelPadron {
			imprimirErrores(new(errores.DNIFueraPadron))
			//continue
		}

		cola.Encolar(TDAVotos.CrearVotante(dni))

		//comando votar
	} else if comandos[0] == "votar" {
		if cola.EstaVacia() {
			imprimirErrores(new(errores.FilaVacia))
			//continue
			//los continue creo hay que sacarlos ya que el for esta en el main, y aca no esta el for
			//golang te recomienda hacer un return, si lo que vos queres hacer es que entre al siguiente ciclo del for, esta bien devolver un return
			//habria q determinar bien que devuelve ese return de todas formas, si nill el erro u otra cosa
		}

		if len(comandos) != 3 {
			imprimirErrores(new(errores.ErrorParametros))
			//continue // si hacemos un return cuando hay errrores, vamos a imprimir el error y hacer que vuelva al inicio del for y pida que escriba correctamente
		}

		var voto TDAVotos.TipoVoto

		//esto habria que ver si funciona o si hay que corregirlo ya que nose si se importo el enum de los votos.
		if comandos[1] == "Presidente" {
			voto = 0
		} else if comandos[1] == "Gobernador" {
			voto = 1
		} else if comandos[1] == "Intendente" {
			voto = 2
		} else {
			imprimirErrores(new(errores.ErrorTipoVoto))
			//continue
		}

		lista, _ := strconv.Atoi(comandos[2])

		// ojo, aca nuestra listaPartidos es un array de partidos tambien.
		//si se hace con arrays, podriamos hacer lista >= len(listaPartidos) y listo
		//forma anterior if lista >= listaPartidos.Largo() {
		if lista >= len(listaPartidos) {
			imprimirErrores(new(errores.ErrorAlternativaInvalida))
			//continue
		}

		_, posicionVotante := esVotanteValido(cola.VerPrimero().LeerDNI(), arrayVotantes)

		errYaVoto := arrayVotantes[posicionVotante].Votar(voto, lista)
		if errYaVoto != nil {
			imprimirErrores(errYaVoto)
		}

		cola.Desencolar()

		//comando deshacer
	} else if comandos[0] == "deshacer" {
		if len(comandos) != 1 {
			imprimirErrores(new(errores.ErrorParametros))
			//continue
		}

		if cola.EstaVacia() {
			imprimirErrores(new(errores.FilaVacia))
			//continue
		}
		errDeshacer := cola.VerPrimero().Deshacer()
		if errDeshacer != nil {
			errVotanteFraudulento := new(errores.ErrorVotanteFraudulento)
			imprimirErrores(errDeshacer)
			//Comparo si el error recibido es ErrorVotanteFraudulento para sacarlo de la fila en cuyo caso
			if errDeshacer.Error() == errVotanteFraudulento.Error() {
				cola.Desencolar()
			}
			//continue
		}

		//comando fin-votar
	} else if comandos[0] == "fin-votar" {
		if len(comandos) != 1 {
			imprimirErrores(new(errores.ErrorParametros))
			//continue
		}

		if cola.EstaVacia() {
			imprimirErrores(new(errores.FilaVacia))
			//continue
		}

		cola.VerPrimero().FinVoto() // supongo aca se guarda el voto final del ingresado ¿?
		//hay que manejar el error que devuevle fin voto en caso de que ya haya votado antes. (leer precondicion de la primitiva fin-voto)
		// creo que hay que guardar esta informacion en el array de votantes, en el votante correspondiente obvio
	}
}

// funcion que recibe un error y lo imprime
func imprimirErrores(err error) {
	fmt.Println(err.Error())
}

//todo 1: definir que hacer con los continue (si se reemplazan por un return o nose que otra opción hay).
// todo 2: definir en que estructura guardamos los partidos
//todo 3: verificar que funcionen las primitivas y corregirlas en caso de que no lo hagan
// todo 4: codear como vamos a imprimir los resultados por stdout (en este script hacemos una funcion que imprima todo, esta funcion debe llamarse desde el main)
