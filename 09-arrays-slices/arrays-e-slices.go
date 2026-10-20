package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Arrays e Slices")

	// Array
	// todos os valroes do array são do mesmo tipo e o tamanho é fixo
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	fmt.Println("Array:", arr)

	array2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Array 2:", array2)

	// Slice
	slice := make([]int, 5)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	slice[4] = 5

	fmt.Println("Slice:", slice)

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice 1:", slice1)
	slice1 = append(slice1, 6)
	fmt.Println("Slice 1 after modification:", slice1)

	fmt.Println("Type of Slice:", reflect.TypeOf(slice))
	fmt.Println("Type of Slice 1:", reflect.TypeOf(slice1))

	fmt.Println("Type of Array 1:", reflect.TypeOf(arr))

	slice2 := array2[1:4]
	fmt.Println("Slice 2:", slice2)
	fmt.Println("Type of Slice 2:", reflect.TypeOf(slice2))

	// ARRAYS INTERNOS
	fmt.Println(" -- ARRAYS INTERNOS -- ")

	slice3 := make([]string, 10, 11)
	fmt.Println("Slice 3:", slice3)
	fmt.Println("Length of Slice 3:", len(slice3))
	fmt.Println("Capacity of Slice 3:", cap(slice3))

	slice3 = append(slice3, "a", "b")
	fmt.Println("Slice 3:", slice3)
	fmt.Println("Length of Slice 3:", len(slice3))
	fmt.Println("Capacity of Slice 3:", cap(slice3))
}
