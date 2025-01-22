package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3} //array -> [3]int
	s1 := []int{1, 2, 3}  //slice -> []int
	fmt.Println(a1, s1)

	a2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Slice não é um array! Slice define um pedaço do array
	s2 := a2[4:9]
	fmt.Println(a2, s2)

	//slice make
	sliceMake := make([]int, 100)
	sliceMake[78] = 12
	fmt.Println(sliceMake)

	sliceMake = make([]int, 10, 20)
	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))
	//slice append
	sliceMake = append(sliceMake, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))

	//Slice usando o mesmo Array
	sliceUsingSameArr := make([]int, 10, 20)
	fmt.Println("Array Criado pelo Make", sliceUsingSameArr)                   // provo  a criação do array
	sliceUsingSameArr2 := append(sliceUsingSameArr, 1, 2, 3, 4, 5, 6, 7, 8, 9) //adicionando os slices dentro do meu array
	fmt.Println("Array com o Slice", sliceUsingSameArr, sliceUsingSameArr2)
	sliceUsingSameArr[0] = 7
	fmt.Println("Array com o Slice1 e 2 provando q usam o mesmo array", sliceUsingSameArr, sliceUsingSameArr2)
}
