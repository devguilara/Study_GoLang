package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2    // se os dois trabalhos derem certo
	comprarTv32 := trab1 != trab2    //ou exclusivo -> se um dos trabalhos derem certo
	comprarSorvete := trab1 || trab2 //se um OU os dois trabalhos derem certo

	return comprarTv50, comprarTv32, comprarSorvete

}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudavel: %t", tv50, tv32, sorvete, !sorvete)
}
