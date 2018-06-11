// SELECTION SORT

/*
Este algoritmo consiste en dividir el arreglo en dos partes
[123,2,3,5,8,81,9]

1: iteramos el arreglo y encontramos el minimo y lo movemos al comienzo

Minimo = 123
Identificamos cual es el menor de la porcion desordenada

Intercambiarlo al principio de la porcion desordenada

[2,123,3,5,8,81,9]

Minimo = 3
[2,3,123,5,8,81,9]...
*/

//Implementacion de Codigo

package main 

import "fmt"

func main(){
	original_array := []int {412,12,123,1,2324,3,155,23412,2}
	fmt.Println(original_array)
	ordered_array := selection_sort(original_array)

	
	fmt.Println(ordered_array)
}

func selection_sort(arreglo []int )[]int{

for i:=0; i< len(arreglo); i++{
	minimo_encontrado,posicion  := arreglo[i],i
	valor_original := arreglo[i]

	// encontrar el minimo en la parte desordenada 

	for j:=i +1; j < len(arreglo); j++{

		valor_comparacion := arreglo[j]


	if valor_comparacion < minimo_encontrado{

		minimo_encontrado, posicion_minimo = valor_comparacion,j
	} 

	}

	if minimo_encontrado != valor_original{

		//intercambiamos posiciones 

		// En Go en una sola linea de codigo
		//arreglo[i],arreglo[posicion_minimo] = minimo_encontrado,valor_original

		//Codigo de forma mas simplificada
		//Intercambio del minimo con el primer elemento del arreglo desordenado
		arreglo[i] = minimo_encontrado

		arreglo[posicion_minimo] = valor_original


	}

	return arreglo

}

}