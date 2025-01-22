package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)

}
func main() {
	result := somar(1, 2)
	imprimir(result)
}
