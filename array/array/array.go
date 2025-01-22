package main

import "fmt"

//Estruturas homogeneas (mesmo tipo) e estática(fixo)

func main() {
	var notas [3]float64
	notas[0], notas[1], notas[2] = 1.4, 8.7, 6.87
	fmt.Printf("notas: %f\n", notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}
	media := total / float64(len(notas))
	fmt.Printf("media: %f\n", media)
}
