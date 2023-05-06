package main

import (
	"bufio"
	"fmt"
	"os"
	TDAOrdenYBusqueda "rerepolez/ordenamiento"
	TDAVotos "rerepolez/votos"
	"strings"

	//"strings"
	//TDACola "tdas/cola"
	"strconv"
)

func main() {

	//todo averiguar como recibir los archivos como parametros asi se leen despues
	padrones, errdni := os.Open(os.Args[1])
	partidos, errPartidos := os.Open(os.Args[2])

	if errdni != nil {
		fmt.Printf("Archivo %s no encontrado\n", padrones)
		return
	}
	defer padrones.Close()

	d := bufio.NewScanner(padrones)
	nroVotantes := 0
	arrayVotantes := []TDAVotos.Votante{}
	//fmt.Println("elementos votantes ", len(arrayVotantes))

	for d.Scan() {
		dni, errAtoi := strconv.Atoi(d.Text())
		votante := TDAVotos.CrearVotante(dni)

		arrayVotantes = append(arrayVotantes, votante)

		nroVotantes++
		if errAtoi != nil {
			fmt.Println(errAtoi)
		}
	}

	errdni = d.Err()
	if errdni != nil {
		fmt.Println(errdni)
	}
	//listo ya funciona el cargar los dnis y ordenarlos:
	//fmt.Println("elementos votantes ", len(arrayVotantes))
	TDAOrdenYBusqueda.OrdenarPadrones(arrayVotantes)
	/*for i := 0; i < nroVotantes; i++ {
		fmt.Println(arrayVotantes[i].LeerDNI())
	}*/

	//todo leer documento partidos y guardar la información como corresponde.
	p := bufio.NewScanner(partidos)

	if errPartidos != nil {
		fmt.Printf("Archivo %s no encontrado\n", partidos)
		return
	}

	defer partidos.Close()

	//creo array de partidos
	arrPartidos := []TDAVotos.Partido{}
	arrPartidos = append(arrPartidos, TDAVotos.CrearVotosEnBlanco())

	for p.Scan() {
		linea := p.Text()
		nombresPartidos := strings.Split(linea, ",")
		nombre := nombresPartidos[0]
		candidatos := [3]string{nombresPartidos[1], nombresPartidos[2], nombresPartidos[3]}
		partidoCreado := TDAVotos.CrearPartido(nombre, candidatos)
		arrPartidos = append(arrPartidos, partidoCreado)
	}

	errPartidos = p.Err()
	if errPartidos != nil {
		fmt.Println(errPartidos)
	}
	for i := 0; i < 4; i++ {
		fmt.Printf("%v\n", arrPartidos[i])
	}
}

//colaIngresantes := TDACola.CrearColaEnlazada[TDAVotos.Votante]()
//scanner := bufio.NewScanner(os.Stdin)

//Aca codigo para leer entrada estandar

// funcion que recibe un error y lo imprime
func imprimirErrores(err error) {
	fmt.Println(err.Error())
}

//Una funcion que me dice si el dni esta dentro del padron o no y me devuelve la posicion dentro del array
/*func esVotanteValido(dni int, votantes []Votante) (bool, int) {
	inicio := 0
	fin := len(votantes) - 1

	for inicio <= fin {
		medio := (inicio + fin) / 2
		if votantes[medio].dni == dni {
			return true, medio
		} else if votantes[medio].dni < dni {
			inicio = medio + 1
		} else {
			fin = medio - 1
		}
	}

	return false, -1

}*/
//aca arranca
//For que termina cuando finaliza entrada estandar
/*for scanner.Scan() {
linea := scanner.Text()
comandos := strings.Split(linea, " ")

if len(comandos) < 1 {
imprimirErrores(new(ErrorParametros))
continue
}*/

//Comando ingresar
/*if comandos[0] == "ingresar" {
if len(comandos) != 2 {
imprimirErrores(new(ErrorParametros))
continue
}
dni, err := strconv.Atoi(comandos[1])
if err != nil {
imprimirErrores(new(DNIError))
continue
}

if dni < 0 {
imprimirErrores(new(DNIError))
continue
}

estaDentroDelPadron,_ := esVotanteValido(dni, arrayVotantes)

if !estaDentroDelPadron {
imprimirErrores(new(DNIFueraPadron))
continue
}

colaIngresantes.Encolar(TDAVotos.CrearVotante(dni))

//comando votar
}else if comandos[0] == "votar" {
if colaIngresantes.EstaVacia() {
imprimirErrores(new(FilaVacia))
continue
}

if len(comandos) != 3 {
imprimirErrores(new(ErrorParametros))
continue
}

var voto TipoVoto

if comandos[1] == "Presidente" {
voto = PRESIDENTE
}else if comandos[1] == "Gobernador" {
voto = GOBERNADOR
}else if comandos[1] == "Intendente" {
voto = INTENDENTE
}else {
imprimirErrores(new(ErrorTipoVoto))
continue
}

lista, _ := strconv.Atoi(comandos[2])

if lista >= listaPartidos.Largo() {
imprimirErrores(new(ErrorAlternativaInvalida))
continue
}

_, posicionVotante := esVotanteValido(colaIngresantes.VerPrimero().LeerDNI(), arrayVotantes)

errYaVoto := arrayVotantes[posicionVotante].Votar(voto)
if errYaVoto != nil {
imprimirErrores(errYaVoto)
}

colaIngresantes.Desencolar()


//comando deshacer
}else if comandos[0] == "deshacer" {
if len(comandos) != 1 {
imprimirErrores(new(ErrorParametros))
continue
}

if colaIngresantes.EstaVacia() {
imprimirErrores(new(FilaVacia))
continue
}
errDeshacer := colaIngresantes.VerPrimero().Deshacer()
if errDeshacer != nil {
errVotanteFraudulento := new(ErrorVotanteFraudulento)
imprimirErrores(errDeshacer)
//Comparo si el error recibido es ErrorVotanteFraudulento para sacarlo de la fila en cuyo caso
if errDeshacer.Error() == errVotanteFraudulento.Error() {
colaIngresantes.Desencolar()
}
continue
}



//comando fin-votar
}else if comandos[0] == "fin-votar" {
if len(comandos) != 1 {
imprimirErrores(new(ErrorParametros))
continue
}

if colaIngresantes.EstaVacia() {
imprimirErrores(new(FilaVacia))
continue
}

colaIngresantes.VerPrimero().FinVoto()
}
}*/

/*onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			return i + 1, data[:i], nil
		}
	}
	if !atEOF {
		return 0, nil, nil
	}
	err = p2.Err()
	if err != nil {
		fmt.Println(err)
	}
	return 0, data, bufio.ErrFinalToken
}*/
