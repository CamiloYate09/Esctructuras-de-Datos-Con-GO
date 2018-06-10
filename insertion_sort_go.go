

//Algoritmo de insercion => Insertion sort


package main 

import "fmt"

func main(){


	arreglo_original := []int{412,12,123,1,2324,3,155,23412,2}
	fmt.Println(arreglo_original)
	arreglo_ordenado := insertion_sort(arreglo_original)
	fmt.Println(arreglo_ordenado)
}

// i = 1
// j = 1
// j -1 = 0
// 412 > 12 

func insertion_sort(arreglo []int) []int{

	for i := 1; i < len(arreglo); i++{
		 j := i 

		 for j > 0 && arreglo[j-1] > arreglo[j] {

		 	//Debug
		 	fmt.Printf("%d se esta intercambiando con %d  \n", arreglo[j-1], arreglo[j])

		 	swap(j-1, j , &arreglo);
		 	j--;
		 }

	}

	return arreglo;


}

// utilizamos un puntero en el tercer parametro 
// los punteros nos indican la posicion en memoria
func swap(previo int , actual int , puntero_arreglo *[]int){

	// actual = 1
	// previo = 0
	// [412,12,123,1,2324,3,155,23412,2]
	//arreglo[previo] = 412
	//arreglo[1] = 412
	//copia = 12
	//arreglo[0] = 12
	arreglo := *puntero_arreglo
	copia := arreglo[actual]
	arreglo[actual] = arreglo[previo]
	arreglo[previo] = copia
}