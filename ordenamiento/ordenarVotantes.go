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
func ComandosLeidos(linea string, cola TDACola.Cola[TDAVotos.Votante], arrayVotantes []TDAVotos.Votante, listaPartidos []TDAVotos.Partido, contadorImpugnados *int) {
	comandos := strings.Split(linea, " ")

	if len(comandos) < 1 {
		return
	}
	// discriminación del comando

	if comandos[0] == "ingresar" {
		if len(comandos) != 2 {
			//imprimirErrores(new(errores.ErrorParametros)) Creo que no hace falta imprimir el error porque creo que es solo para los archivos este error)
			return
		}
		dni, err := strconv.Atoi(comandos[1])
		if err != nil {
			imprimirErrores(new(errores.DNIError))
			return
		}

		if dni <= 0 {
			imprimirErrores(new(errores.DNIError))
			return
		}

		estaDentroDelPadron, _ := esVotanteValido(dni, arrayVotantes)

		if !estaDentroDelPadron {
			imprimirErrores(new(errores.DNIFueraPadron))
			return
		}

		cola.Encolar(TDAVotos.CrearVotante(dni))
		fmt.Println("OK")

		//comando votar
	} else if comandos[0] == "votar" {
		if len(comandos) != 3 {
			//imprimirErrores(new(errores.ErrorParametros))
			return // si hacemos un return cuando hay errrores, vamos a imprimir el error y hacer que vuelva al inicio del for y pida que escriba correctamente
		}

		if cola.EstaVacia() {
			imprimirErrores(new(errores.FilaVacia))
			return
			//los continue creo hay que sacarlos ya que el for esta en el main, y aca no esta el for
			//golang te recomienda hacer un return, si lo que vos queres hacer es que entre al siguiente ciclo del for, esta bien devolver un return
			//habria q determinar bien que devuelve ese return de todas formas, si nill el erro u otra cosa
		}

		var voto TDAVotos.TipoVoto

		//esto habria que ver si funciona o si hay que corregirlo ya que nose si se importo el enum de los votos.
		//Yo supongo que si se importa, hay que cambiar eso
		if comandos[1] == "Presidente" {
			voto = TDAVotos.PRESIDENTE
		} else if comandos[1] == "Gobernador" {
			voto = TDAVotos.GOBERNADOR
		} else if comandos[1] == "Intendente" {
			voto = TDAVotos.INTENDENTE
		} else {
			imprimirErrores(new(errores.ErrorTipoVoto))
			return
		}

		lista, errorlista := strconv.Atoi(comandos[2])
		if errorlista != nil {
			imprimirErrores(new(errores.ErrorAlternativaInvalida))
			return
		}
		// ojo, aca nuestra listaPartidos es un array de partidos tambien.
		//si se hace con arrays, podriamos hacer lista >= len(listaPartidos) y listo
		//forma anterior if lista >= listaPartidos.Largo() {
		if lista >= len(listaPartidos) {
			imprimirErrores(new(errores.ErrorAlternativaInvalida))
			return
		}

		_, posicionVotante := esVotanteValido(cola.VerPrimero().LeerDNI(), arrayVotantes)

		errYaVoto := arrayVotantes[posicionVotante].Votar(voto, lista)
		if errYaVoto != nil {
			imprimirErrores(errYaVoto)
			cola.Desencolar()
			return
		}
		fmt.Println("OK")

		//comando deshacer
	} else if comandos[0] == "deshacer" {
		if len(comandos) != 1 {
			//imprimirErrores(new(errores.ErrorParametros))
			return
		}

		if cola.EstaVacia() {
			imprimirErrores(new(errores.FilaVacia))
			return
		}
		_, posicionVotante := esVotanteValido(cola.VerPrimero().LeerDNI(), arrayVotantes)

		errDeshacer := arrayVotantes[posicionVotante].Deshacer()

		if errDeshacer != nil {
			imprimirErrores(errDeshacer)
			//Comparo si el error recibido es ErrorVotanteFraudulento para sacarlo de la fila en cuyo caso
			if errDeshacer == (errores.ErrorVotanteFraudulento{Dni: cola.VerPrimero().LeerDNI()}) {
				cola.Desencolar()
			}
			return
		}
		fmt.Println("OK")
		//comando fin-votar
	} else if comandos[0] == "fin-votar" {
		if len(comandos) != 1 {
			//imprimirErrores(new(errores.ErrorParametros))
			return
		}

		if cola.EstaVacia() {
			imprimirErrores(new(errores.FilaVacia))
			return
		}
		_, posicionVotante := esVotanteValido(cola.VerPrimero().LeerDNI(), arrayVotantes)
		voto, errFinVoto := arrayVotantes[posicionVotante].FinVoto()
		if errFinVoto != nil {
			imprimirErrores(errFinVoto)
			cola.Desencolar()
			return
		}
		cola.Desencolar()
		if voto.Impugnado {
			*contadorImpugnados = *contadorImpugnados + 1
			fmt.Println("OK")
			return
		}
		for i := 0; i < 3; i++ {
			if i == 0 {
				listaPartidos[voto.VotoPorTipo[i]].VotadoPara(TDAVotos.PRESIDENTE)
			} else if i == 1 {
				listaPartidos[voto.VotoPorTipo[i]].VotadoPara(TDAVotos.GOBERNADOR)
			} else {
				listaPartidos[voto.VotoPorTipo[i]].VotadoPara(TDAVotos.INTENDENTE)
			}

		}
		fmt.Println("OK")
		// supongo aca se guarda el voto final del ingresado ¿?
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
