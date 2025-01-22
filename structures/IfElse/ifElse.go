package main

import (
	"fmt"
	"math/rand"
	"time"
)

func imprimeResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado!!! :)")
	} else {
		fmt.Println("Reprovado!!! :(")
	}
}

func notas(nota float64) {
	if nota >= 9 && nota <= 10 {
		fmt.Println("Aprovado com média alta!!! :)")
	} else if nota >= 7 && nota <= 9 {
		fmt.Println("Aprovado com média normal")
	} else if nota >= 6 && nota <= 7 {
		fmt.Println("Aprovado com média normal, mas da pra melhorar")
	} else if nota >= 4 && nota <= 6 {
		fmt.Println("Reprovado com chances de se recuperar fazendo exame")
	} else {
		fmt.Println("Reprovado sem chances de recuperação")
	}
}

//desafio transformar func notas em switch

// bloco de inicialização de variável
func IFInit() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}
func main() {
	imprimeResultado(7)
	imprimeResultado(2)

	notas(7)
	notas(9.7)
	notas(2.4)
	notas(7.77)
	notas(6.95)

	if i := IFInit(); i > 5 {
		fmt.Println("Ganhou", i)
	} else {
		fmt.Println("Perdeu", i)
	}

}
