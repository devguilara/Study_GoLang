package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 // float
	y := 2   // inteiro

	fmt.Println(x / float64(y)) //converter explicitamente o tipo inteiro para float64

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println("Valor da nota arredondada  ", notaFinal)

	//******CUIDADO******

	fmt.Println("Teste" + string(187)) //string(187) é referente ao valor int32 na tabela unicode e nao o valor 187

	// int para string corretamente

	fmt.Println("Teste" + strconv.Itoa(123))

	//string para int
	num, _ := strconv.Atoi("123") //tem q ser passado 2 valores -> variavel(num) e um exception caso nao seja  um inteiro valido (_)
	fmt.Println(num)

	//string para boleano
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("true")
	}

}
