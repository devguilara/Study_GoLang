package main

import "fmt"

func main() {

	i := 1

	for i <= 10 { // nao temos WHILE em Go
		fmt.Println(i)
		i++
	}

	for x := 1; x <= 10; x++ {
		fmt.Println(x)
	}

	for y := 1; y <= 10; y += 2 {
		fmt.Println(y)
	}

	fmt.Println("Misturando")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}
}
