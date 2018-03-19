package main

import "fmt"

func main() {
	arreglo :=[] int{1};
	fmt.Println(arreglo)
	arreglo= append(arreglo,7)
	fmt.Println(arreglo)
	arreglo= append(arreglo[:0], arreglo[0+1:]...)
	fmt.Println(arreglo)

}

