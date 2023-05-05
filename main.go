package main

import (
	"bufio"
	"fmt"
	"os"
	TDAOrdenYBusqueda "rerepolez/ordenamiento"
	TDAVotos "rerepolez/votos"
	//TDAPila "tdas/pila"
	"strconv"
)

func main() {
	//todo averiguar como recibir los archivos como parametros asi se leen despues
	padrones, errdni := os.Open("tests/02_padron")

	if errdni != nil {
		fmt.Printf("Archivo %s no encontrado\n", padrones)
		return
	}
	defer padrones.Close()

	d := bufio.NewScanner(padrones)
	nroVotantes := 0
	arrayVotantes := []TDAVotos.Votante{}
	fmt.Println("elementos votantes ", len(arrayVotantes))

	for d.Scan() {
		dni, errAtoi := strconv.Atoi(d.Text())
		votante := TDAVotos.CrearVotante(dni)
		fmt.Printf("Leí: %v\n", dni)
		//fmt.Printf("De la fila: %v \n", nroVotantes)
		arrayVotantes = append(arrayVotantes, votante)
		//pilaVotantes.Apilar(votante)
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
	fmt.Println("elementos votantes ", len(arrayVotantes))
	TDAOrdenYBusqueda.OrdenarPadrones(arrayVotantes)
	for i := 0; i < nroVotantes; i++ {
		fmt.Println(arrayVotantes[i].LeerDNI())
	}

	//todo leer documento partidos y guardar la información como corresponde.
	partidos, errPartidos := os.Open("tests/02_partidos")
	p := bufio.NewScanner(partidos)
	p2 := bufio.NewScanner(partidos)

	if errPartidos != nil {
		fmt.Printf("Archivo %s no encontrado\n", partidos)
		return
	}

	defer partidos.Close()

	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
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
	}

	p2.Split(onComma)

	columna := 0

	for p2.Scan() {
		fmt.Printf("De la columna: %v \n", columna)
		fmt.Printf("Leí: %s\n", p2.Text())
		columna++
		if columna > 3 {
			columna = 0
		}
	}
	fmt.Println("salto de linea")

	columna = 0
	errPartidos = p.Err()
	if errPartidos != nil {
		fmt.Println(errPartidos)
	}
	
	    //Aca codigo para leer entrada estandar

    //funcion que recibe un error y lo imprime
    func imprimirErrores(err error) {
        fmt.Println(err.Error())
    }
    
    //Una funcion que me dice si el dni esta dentro del padron o no y me devuelve la posicion dentro del array
    func esVotanteValido(dni int, votantes []Votante) (bool, int) {
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
    }

    //aca arranca

    colaIngresantes := TDACola.CrearColaEnlazada[Votante]()

    scanner := bufio.NewScanner(os.Stdin)

    //For que termina cuando finaliza entrada estandar
    for scanner.Scan() {
        linea := scanner.Text()
        comandos := strings.Split(linea, " ")

        if len(comandos) < 1 {
            imprimirErrores(new(ErrorParametros))
            continue
        }

        //Comando ingresar
        if comandos[0] == "ingresar" {
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
    }
}
