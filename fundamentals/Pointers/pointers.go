package main

import "fmt"

func main() {
	// Go nao tem aritimética de ponteiros
	i := 1 // endereço com equivalencia de 64 bits ou 8 bytes

	//nil = nulo
	var p *int = nil

	p = &i // p agora tem o mesmo endereço de memoria apontado para a variavel i
	*p++   // desreferenciando o ponteiro (pegando o valor)
	i++
	fmt.Println(p, *p, i)
}
