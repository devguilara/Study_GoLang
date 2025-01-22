package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2
	//Forma reduzida de iniciar variavel -> variavel1, variavel2, n.... := <valor1,valor2,valorN...>
	area := PI * math.Pow(raio, 2)
	fmt.Printf("%.2f\n", area)
}
