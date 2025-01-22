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
	sliceMake = append(sliceMake, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))
}
