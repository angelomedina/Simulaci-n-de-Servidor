package main


import "C"

import (
	"math/rand"
	"time"
	"fmt"
	"sync"
)

// sincronizacion de rutinas
var wg sync.WaitGroup

// las rutas o puentes de acceso a la rotondo
//export colaA
var colaA = []auto  {}
//export colaB
var colaB = []auto  {}
//export colaC
var colaC = []auto  {}

//export rotonda
//rotonda solo permite 6 vehiculos a la vez
var rotonda =[] auto{}

//export auto
//structura de ruta de auto
type auto struct {
	origen string
	destino string
	duracion time.Duration
	}

//export cargaPaquetesAutos
//agrega autos a sus rutas
func cargaPaquetesAutos(origen string, tamaño int)  {
	/*
		recibo un string de origen y un tamaño( cantidad de autos que vienn de esa ruta)
	*/

	//numero aleatorio para escoger   el destino
	 rand.Seed(time.Now().UnixNano())
	 //escoge entre esos destinos
	 rutas := [3]string{"Florecia","Santa Clara","San Ramón"}

	 var i int
	 for i=0 ; i < tamaño; i++ {
	 	// random para escger el destino
		randomNum := random(0, 3)
		destino   := rutas[randomNum ]

		//el ot}rigen y el destino no puden seriguales
		if equal(origen, destino) == 0{
			//se crea unnuevo auto
			nuevo := auto{origen,destino,0}
			//origen y destino disstintos
			if equal(nuevo.origen,"Florecia") == 0{
				colaA = append(colaA,auto{origen:nuevo.origen,destino:nuevo.destino})
			}
			if equal(nuevo.origen,"Santa Clara") == 0{
				colaB = append(colaB,auto{origen:nuevo.origen,destino:nuevo.destino})

			}
			if equal(nuevo.origen,"San Ramón") == 0{
				colaC = append(colaC,auto{origen:nuevo.origen,destino:nuevo.destino})
			}
		}
	 }
}

//export caragarRutas
//carga las rutas con sus eventuales destinos
func caragarRutas() {
	//random
	rand.Seed(time.Now().UnixNano())
	//rutas origen
	rutas := [3]string{"Florecia","Santa Clara","San Ramón"}
	//pquetes de utos
	paquetesAutos := [] int{6,5,4,3,2,1,0}

	//carga rutas de A por defecto
	var i int
	for i=0 ; i < len(rutas); i++ {
		origen          := rutas[i]
		randomTamaño    := random(0, 5)
		tamaño          := paquetesAutos[randomTamaño]

		//lo que hace es cargar filas o grupos  de autos
		println("Ruta: ",origen ,"tiene en fila ",tamaño,"vehiclos")
		cargaPaquetesAutos(origen , tamaño)
	}
}

//export random
//genera el random del destino de la ruta origen
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

//export equal
//funcion para igualar strings
func equal(s1, s2 string) int {
	eq := 0
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	for key, _ := range s1 {
		if s1[key] == s2[key] {
			eq++
		} else {
			break
		}
	}
	return eq
}


//export servidor
//simula la carga vehiclar el la rotonda
func servidor()  {

	wg.Add(3)

	fmt.Println("Iniciamos las goruntinas")

	go solicitudCola(colaA,lenghtCola(colaA))
	go solicitudCola(colaB,lenghtCola(colaB))
	go solicitudCola(colaC,lenghtCola(colaC))

	wg.Wait()
	fmt.Println("\nTerminando el programa")

}

//export lenghtCola
//retorna la cantidad de autos en cola  de cad ruta
func lenghtCola(cola []auto)int  {
	a:=0
	for _, recorre := range cola {
		recorre.origen=recorre.origen
		a++
	}
	return a
}

//export solicitudCola
//ejecula la cola delas rutas/solicitud de acceso
func solicitudCola(cola []auto, tamaño int)  {
	//ejecuta la llamada de gorutina
	defer wg.Done()

		//llama solo si la ruta tiene mas de 0 carros
		if tamaño > 0 {
			//hago la cola de la ruta
			for cantidad := 1; cantidad <= tamaño; cantidad ++ {

				//impresion de solicitud de inicio
				fmt.Println("\n----->SOLICIDUD DE ACCESO:", cola[0].origen+"-->"+cola[0].destino)
				/*
					la rotonda solo pueden tener 6 veiculos a la vez, pero
					como los vehiculos salen de la rotonda, se van tanto agregando como saliendo vehiculos de la rotonda
				*/
				if( lenghtCola(rotonda) < 5 ){
					//tiempo de espera de permitido
					sleep := rand.Int63n(3)
					time.Sleep(time.Duration(sleep) * time.Second)

					//impresion de solicitud de permisos
					fmt.Println("\n[A][D][E][L][A][N][TE] Acceso permitido:", cola[0].origen+"-->"+cola[0].destino)

					//agrega auto a la rotonda
					rotonda = append(rotonda,auto{cola[0].origen,cola[0].destino,0})
					//ejecuta la cola de la rotonda
					respuestaCola(cola)

				}else{
					//tiempo de espera de espera
					sleep := rand.Int63n(3)
					time.Sleep(time.Duration(sleep) * time.Second)

					fmt.Println("\nAcceso espera:", cola[0].origen+"-->"+cola[0].destino)
				}
			}
		}
}

//export  respuestaCola
//ejecuta el acceso de la solicitud
func respuestaCola(cola []auto)  {


	//doy tiempo de impresion
	sleep := rand.Int63n(3)
	time.Sleep(time.Duration(sleep) * time.Second)

	if equal(cola[0].origen, "Florecia") == 0 {
		cola[0].duracion=(time.Duration(sleep)+1) * time.Minute
		fmt.Println("\n<----SALIDA DE PUENTE",cola[0].origen+"-->"+cola[0].destino,cola[0].duracion)
		//elimono el elemento de la cola dela ruta

		if (lenghtCola(colaA) >1) {
			colaA = append(colaA[:0], colaA[0+1:]...)
		}else{
			//deja sin elementos la lista
			colaA=colaA[:0]
		}
	}
	if equal(cola[0].origen, "Santa Clara") == 0 {
		cola[0].duracion=(time.Duration(sleep)+1) * time.Minute
		fmt.Println("\n<----SALIDA DE ROTONDA",cola[0].origen+"-->"+cola[0].destino,cola[0].duracion)
		//elimono el elemento de la cola dela ruta

		if (lenghtCola(colaB) >1) {
			colaB = append(colaB[:0], colaB[0+1:]...)
		}else{
			//deja sin elementos la lista
			colaB=colaB[:0]
		}
	}
	if equal(cola[0].origen, "San Ramón") == 0 {

		cola[0].duracion=(time.Duration(sleep)+1) * time.Minute
		fmt.Println("\n<----SALIDA DE ROTONDA",cola[0].origen+"-->"+cola[0].destino,cola[0].duracion)
		//elimono el elemento de la cola dela ruta

		if (lenghtCola(colaC) >1) {
			colaC = append(colaC[:0], colaC[0+1:]...)
		}else{
			//deja sin elementos la lista
			colaC=colaC[:0]
		}

	}

	//autos saliendo de la rotonda y eleimina el vehiculo de la rotonda
	rotonda= append(rotonda[:0],rotonda[0+1:]...)
}


func main() {
	for i:= 1; i<=3;i++ {

		caragarRutas();
		servidor();
	}
}
