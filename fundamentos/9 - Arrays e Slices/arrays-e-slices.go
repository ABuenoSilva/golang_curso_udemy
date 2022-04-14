package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	//ARRAYS
	//Arrays não podem ser dinâmicos, a qtde de elementos sempre é a mesma
	var array1 [5]int
	var array2 [5]string
	fmt.Println(array1, array2)

	array3 := [5]string{"Pos1", "Pos2", "Pos3", "Pos4", "Pos5"}
	fmt.Println(array3)

	//Isto não muda se é dinâmico ou não
	array4 := [...]int{1, 2, 3, 4}
	fmt.Println(array4)

	//SLICES
	slice := []int{10, 12, 14, 16}
	fmt.Println(slice, slice[1], reflect.TypeOf(slice), reflect.TypeOf(array4))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array3[1:3]
	fmt.Println(slice2)
	array3[1] = "Alterei"
	fmt.Println(slice2)

	//ARRAYS INTERNOS
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3, len(slice3), cap(slice3))

	slice3 = append(slice3, 3)
	fmt.Println(slice3, len(slice3), cap(slice3))

	slice3 = append(slice3, 5)
	fmt.Println(slice3, len(slice3), cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4, len(slice4), cap(slice4))
	slice4 = append(slice4, 10)
	fmt.Println(slice4, len(slice4), cap(slice4))
}
