package main

import (
	"fmt"
	"time"
)

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3, 2:
		return "D"
	case 1:
		return "E"
	default:
		return "Nota inválida"
	}

}
func main() {
	fmt.Println(notaParaConceito(10.0))
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	case t.Hour() < 19:
		fmt.Println("Good evening!")
	default:
		fmt.Println("Good night!")
	}
}
