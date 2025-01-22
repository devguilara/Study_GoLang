package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma => ", a+b)
	fmt.Println("Sub => ", a-b)
	fmt.Println("Div => ", a/b)
	fmt.Println("Multi => ", a*b)
	fmt.Println("Modulo(Resto da Divisão) => ", a%b)

	//bitwise
	fmt.Println("E => ", a&b)    // 11 & 10 = 10
	fmt.Println("OU => ", a|b)   // 11 | 10 = 11
	fmt.Println("XOR => ", a^b)  // 11 ^ 10 = 01
	fmt.Println("AND => ", a&^b) //11 &^ 10 = 01

	// outros operações usando math

	fmt.Println("Maior => ", math.Max(float64(a), float64(b)))
	fmt.Println("Menor => ", math.Min(float64(a), float64(b)))
	fmt.Println("Exponencial => ", math.Pow(float64(a), float64(b)))

}
