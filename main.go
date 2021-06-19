package main

import "fmt"

func main() {

	var objetivo int8

	slice := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("Tenemos el slice: %v\n", slice)
	fmt.Printf("Escriba el numero que desea buscar: ")
	fmt.Scan(&objetivo)

	fmt.Println()
	fmt.Println("Con busqueda recursiva: ")
	index := busquedaBinariaRecursiva(slice, 0, int8(len(slice)-1), objetivo)

	if index == -1 {
		fmt.Println("El numero no fue encontrado")
	} else {
		fmt.Printf("El numero se encontro en el indice: %d\n", index)
	}

	fmt.Println()
	fmt.Println("Con busqueda iterativa: ")

	index = busquedaBinariaIterativa(slice, 0, int8(len(slice)-1), objetivo)

	if index == -1 {
		fmt.Println("El numero no fue encontrado")
	} else {
		fmt.Printf("El numero se encontro en el indice: %d\n", index)
	}
}

// siendo funcion O(n^2)
func busquedaBinariaRecursiva(slice []int8, left, right, objetivo int8) int8 {

	if right >= left {
		var mid = left + (right-left)/2

		if slice[mid] == objetivo {
			return mid
		}

		if slice[mid] > objetivo {
			return busquedaBinariaRecursiva(slice, left, mid-1, objetivo)
		}
		return busquedaBinariaRecursiva(slice, mid+1, right, objetivo)
	}

	return -1
}

// siendo funcion O(n log2)
func busquedaBinariaIterativa(slice []int8, left, right, objetivo int8) int8 {

	for left <= right {
		var mid = left + (right-left)/2

		if slice[mid] == objetivo {
			return mid
		}

		if slice[mid] < objetivo {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
